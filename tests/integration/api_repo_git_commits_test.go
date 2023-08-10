// SPDX-License-Identifier: MIT
	auth_model "code.gitea.io/gitea/models/auth"
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	MakeRequest(t, req, http.StatusNotFound)
	MakeRequest(t, req, http.StatusUnprocessableEntity)
	MakeRequest(t, req, http.StatusNotFound)
		MakeRequest(t, req, http.StatusOK)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)

	// Test getting commits (Page 1)
	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo20/commits?token="+token+"&not=master&sha=remove-files-a", user.Name)
	resp := MakeRequest(t, req, http.StatusOK)

	var apiData []api.Commit
	DecodeJSON(t, resp, &apiData)

	assert.Len(t, apiData, 2)
	assert.EqualValues(t, "cfe3b3c1fd36fba04f9183287b106497e1afe986", apiData[0].CommitMeta.SHA)
	compareCommitFiles(t, []string{"link_hi", "test.csv"}, apiData[0].Files)
	assert.EqualValues(t, "c8e31bc7688741a5287fcde4fbb8fc129ca07027", apiData[1].CommitMeta.SHA)
	compareCommitFiles(t, []string{"test.csv"}, apiData[1].Files)

	assert.EqualValues(t, resp.Header().Get("X-Total"), "2")
}

func TestAPIReposGitCommitListNotMaster(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2})
	// Login as User2.
	session := loginUser(t, user.Name)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	resp := MakeRequest(t, req, http.StatusOK)

	assert.EqualValues(t, resp.Header().Get("X-Total"), "3")
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	resp := MakeRequest(t, req, http.StatusOK)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	resp := MakeRequest(t, req, http.StatusOK)
func TestAPIReposGitCommitListWithoutSelectFields(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2})
	// Login as User2.
	session := loginUser(t, user.Name)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)

	// Test getting commits without files, verification, and stats
	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo16/commits?token="+token+"&sha=good-sign&stat=false&files=false&verification=false", user.Name)
	resp := MakeRequest(t, req, http.StatusOK)

	var apiData []api.Commit
	DecodeJSON(t, resp, &apiData)

	assert.Len(t, apiData, 1)
	assert.Equal(t, "f27c2b2b03dcab38beaf89b0ab4ff61f6de63441", apiData[0].CommitMeta.SHA)
	assert.Equal(t, (*api.CommitStats)(nil), apiData[0].Stats)
	assert.Equal(t, (*api.PayloadCommitVerification)(nil), apiData[0].RepoCommit.Verification)
	assert.Equal(t, ([]*api.CommitAffectedFiles)(nil), apiData[0].Files)
}

	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	resp := MakeRequest(t, reqDiff, http.StatusOK)
	resp = MakeRequest(t, reqPatch, http.StatusOK)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)
	resp := MakeRequest(t, req, http.StatusOK)

	assert.EqualValues(t, resp.Header().Get("X-Total"), "1")
}

func TestGetFileHistoryNotOnMaster(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2})
	// Login as User2.
	session := loginUser(t, user.Name)
	token := getTokenForLoggedInUser(t, session, auth_model.AccessTokenScopeReadRepository)

	req := NewRequestf(t, "GET", "/api/v1/repos/%s/repo20/commits?path=test.csv&token="+token+"&sha=add-csv&not=master", user.Name)
	resp := MakeRequest(t, req, http.StatusOK)

	var apiData []api.Commit
	DecodeJSON(t, resp, &apiData)

	assert.Len(t, apiData, 1)
	assert.Equal(t, "c8e31bc7688741a5287fcde4fbb8fc129ca07027", apiData[0].CommitMeta.SHA)
	compareCommitFiles(t, []string{"test.csv"}, apiData[0].Files)

	assert.EqualValues(t, resp.Header().Get("X-Total"), "1")