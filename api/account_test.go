package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	mockdb "github.com/sinazrp/golang-bank/db/mock"
	db "github.com/sinazrp/golang-bank/db/sqlc"
	"github.com/sinazrp/golang-bank/util"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccountAPI(t *testing.T) {

	// create a random account
	account := randomAccount()

	testCases := []struct {
		name          string
		accountID     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "NotFound",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
				requiredBodyMatchAccount(t, recorder.Body, db.Account{})
			},
		},
		{
			name:      "OK",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requiredBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "internalError",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
				requiredBodyMatchAccount(t, recorder.Body, db.Account{})
			},
		},
		{
			name:      "invalidID",
			accountID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			// create a mock controller
			ctrl := gomock.NewController(t)
			// when the test is finished, finish the mock controller
			defer ctrl.Finish()

			// create a mock store
			store := mockdb.NewMockStore(ctrl)

			// build stubs
			tc.buildStubs(store)

			// create a new server with the mock store
			server := NewTestServer(store)

			// create a new recorder
			recorder := httptest.NewRecorder()

			// create a new request with the GET method and the url /accounts/{account.ID}
			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// call the server with the request and recorder
			server.router.ServeHTTP(recorder, request)

			// check the response
			tc.checkResponse(t, recorder)

		})

	}

}

func TestCreateAccountAPI(t *testing.T) {

	// create a random account
	account := randomAccount()
	account.Balance = 0

	testCases := []struct {
		name string
		body struct {
			Owner    string
			Currency string
		}
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: struct {
				Owner    string
				Currency string
			}{
				Owner:    account.Owner,
				Currency: account.Currency,
			},

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateAccount(gomock.Any(), gomock.Eq(db.CreateAccountParams{
					Owner:    account.Owner,
					Balance:  0,
					Currency: account.Currency})).Times(1).Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requiredBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name: "invalidCurrency",
			body: struct {
				Owner    string
				Currency string
			}{
				Owner:    account.Owner,
				Currency: "Rials",
			},

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
		{
			name: "internalError",
			body: struct {
				Owner    string
				Currency string
			}{
				Owner:    account.Owner,
				Currency: account.Currency,
			},

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Times(1).Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			// create a mock controller
			ctrl := gomock.NewController(t)
			// when the test is finished, finish the mock controller
			defer ctrl.Finish()

			// create a mock store
			store := mockdb.NewMockStore(ctrl)

			// build stubs
			tc.buildStubs(store)

			// create a new server with the mock store
			server := NewTestServer(store)

			// create a new recorder
			recorder := httptest.NewRecorder()
			jsonBody, err := json.Marshal(tc.body)
			if err != nil {
				t.Fatal(err)
			}

			// create a new request with the GET method and the url /accounts/{account.ID}
			url := fmt.Sprintf("/accounts")
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))

			require.NoError(t, err)

			// call the server with the request and recorder
			server.router.ServeHTTP(recorder, request)

			// check the response
			tc.checkResponse(t, recorder)

		})
	}
}
func randomAccount() db.Account {

	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requiredBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)
	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)

}
