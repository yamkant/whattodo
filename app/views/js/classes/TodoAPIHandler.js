class TodoAPIHandler {
    static updateTodoCompleted = async (target) => {
        const todoItemObj = target.closest('.todo-item')
        const checkboxInputObj = todoItemObj.querySelector('input[type="checkbox"]')

        const data = { completed: checkboxInputObj.checked }

        const item_id = target.closest('li').dataset.value;
        const resData = await axios({
            method: 'patch',
            url: `/api/v1/todos/${item_id}/`,
            headers: {
                'Content-Type': 'application/json'
            },
            data: data,
        }).then((response) => {
            return getAxiosResponseData(response);
        }).catch((err) => {
            console.error(err);
        });
        if (resData) {
            setCheckBoxByChecked(checkboxInputObj);
        }
    }

    static updateTodoContent = async (target) => {
        const todoItemObj = target.closest('.todo-item')
        const todoContentInputObj = todoItemObj.querySelector('.todo-content');
        const todoContent= todoContentInputObj.value;

        const data = { content: todoContent }
        const item_id = target.closest('li').dataset.value;
        const resData = await axios({
            method: 'patch',
            url: `/api/v1/todos/${item_id}/content/`,
            headers: {
                'Content-Type': 'application/json'
            },
            data: data,
        }).then((response) => {
            return getAxiosResponseData(response);
        }).catch((err) => {
            console.error(err);
        });
    }

    static updateTodoStartAt = async (target) => {
        const item_id = target.closest('li').dataset.value;
        const resData = await axios({
            method: 'patch',
            url: `/api/v1/todos/${item_id}/start_at/`,
            headers: {
                'Content-Type': 'application/json'
            },
        }).then((response) => {
            return getAxiosResponseData(response);
        }).catch((err) => {
            console.error(err);
        });

        if (!resData) {
            alert('Server error');
            return;
        }

        location.reload();
    }

    static updateTodoEndAt = async (target) => {
        const item_id = target.closest('li').dataset.value;
        const resData = await axios({
            method: 'patch',
            url: `/api/v1/todos/${item_id}/end_at/`,
            headers: {
                'Content-Type': 'application/json'
            },
        }).then((response) => {
            return getAxiosResponseData(response);
        }).catch((err) => {
            console.error(err);
        });

        if (!resData) {
            alert('Push start button first.');
            return;
        }

        location.reload();
    }
}
