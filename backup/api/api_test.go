package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"example.com/m/model"
	"github.com/stretchr/testify/assert"
)

func getTodoBodyData(resp *http.Response) (model.Todo, error) {
	var todo model.Todo
	err := json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func getTodoBodyDataList(resp *http.Response) ([]*model.Todo, error) {
	todos := []*model.Todo{}
	err := json.NewDecoder(resp.Body).Decode(&todos)
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func getDeleteAPIResponseData(resp *http.Response) (Success, error) {
	var success Success
	err := json.NewDecoder(resp.Body).Decode(&success)
	if err != nil {
		return success, err
	}
	return success, nil
}

func TestTodos(t *testing.T) {
	assert := assert.New(t)
	ah := APIHttpHandler()

	ts := httptest.NewServer(ah)
	ts_url := ts.URL + "/api/todos/"
	defer ts.Close()

	// NOTE: CREATE TEST
	for i := 0; i < 2; i++ {
		todoContent := "Test todo " + strconv.Itoa(i)
		tbytes, _ := json.Marshal(AddTodoDTO{todoContent})
		resp, err := http.Post(ts_url, "application/json", bytes.NewBuffer(tbytes))
		assert.NoError(err)
		assert.Equal(http.StatusCreated, resp.StatusCode)

		todo, err := getTodoBodyData(resp)
		assert.NoError(err)
		assert.Equal(todoContent, todo.Content)
	}

	// NOTE: GET TEST
	resp, err := http.Get(ts_url)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos, err := getTodoBodyDataList(resp)
	assert.NoError(err)
	assert.Equal(len(todos), 2)

	// NOTE: PATCH TEST
	pbytes, _ := json.Marshal(UpdateTodoDTO{time.Time{}, time.Time{}, true})
	// patchData := []byte(`{"completed": true}`)
	req, _ := http.NewRequest("PATCH", ts_url+strconv.Itoa(todos[0].ID)+"/", bytes.NewBuffer(pbytes))
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts_url)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos, err = getTodoBodyDataList(resp)
	assert.NoError(err)
	for _, todo := range todos {
		if todo.ID == todos[0].ID {
			assert.True(todo.Completed)
		}
	}

	// NOTE: DELETE TEST
	req, _ = http.NewRequest("DELETE", ts_url+strconv.Itoa(todos[0].ID)+"/", nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	result, err := getDeleteAPIResponseData(resp)
	assert.NoError(err)
	assert.True(result.Flag)

	resp, err = http.Get(ts_url)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos, err = getTodoBodyDataList(resp)
	assert.NoError(err)
	assert.Equal(len(todos), 1)
}
