function getAxiosResponseData(response) {
    return response.data.data;
}

const getCheckboxLabelValue = (checked) => {
    return checked ? `
        <svg class="w-6 h-6 mr-2 text-green-500 dark:text-green-400 flex-shrink-0" aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
            <path
                d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5Zm3.707 8.207-4 4a1 1 0 0 1-1.414 0l-2-2a1 1 0 0 1 1.414-1.414L9 10.586l3.293-3.293a1 1 0 0 1 1.414 1.414Z" />
        </svg>
    ` : `
        <svg class="w-6 h-6 mr-2 text-gray-500 dark:text-gray-400 flex-shrink-0" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
            <path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5Zm3.707 8.207-4 4a1 1 0 0 1-1.414 0l-2-2a1 1 0 0 1 1.414-1.414L9 10.586l3.293-3.293a1 1 0 0 1 1.414 1.414Z"/>
        </svg>
    `;
}

const displayDateTime = (dateTime) => {
    return dateTime.slice(11, 16) === '00:00' ? "" : dateTime.slice(11, 16);
}

const onClickRemoveTodo = async (target) => {
    const listObj = target.closest('li');
    const itemId = listObj.dataset.value;
    if (confirm("Are you sure you want to remove?")) {
        const res = await axios({
            method: 'delete',
            url: `/api/v1/todos/${itemId}/`,
        })
        .then((response) => {
            return getAxiosResponseData(response);
        })
        .catch((err) => {
            console.error(err);
        });
        if (res) {
            listObj.remove();
        }
    }
    location.reload();
}

const setCheckBoxByChecked = (target) => {
    const targetLabel = target.closest('.checkbox-container').querySelector('label');
    targetLabel.innerHTML = getCheckboxLabelValue(target.checked);
}

const onClickAddBtn = async (target) => {
    const todoUlContainer = document.querySelector('.todo-ul-container');
    const todoUl = todoUlContainer.querySelector('.todo-ul-uncompleted');

    const inputObj = target.closest('.todo-form').querySelector('input');

    const resData = await axios({
        method: 'post',
        url: '/api/v1/todos/',
        headers: {
            'Content-Type': 'application/json'
        },
        data: {
            content: inputObj.value,
        }
    })
    .then((response) => {
        return getAxiosResponseData(response);
    })
    .catch((err) => {
        console.error(err);
    });
    location.reload();
}


window.addEventListener('load', async () => {
    const todoUlContainer = document.querySelector('.todo-ul-container');
    const todoUlUncompleted = todoUlContainer.querySelector('.todo-ul-uncompleted');
    const todoUlCompleted = todoUlContainer.querySelector('.todo-ul-completed');

    const resData = await axios({
        method: 'get',
        url: '/api/v1/todos/',
    })
    .then((response) => {
        return getAxiosResponseData(response);
    })
    .catch((err) => {
        console.error(err);
    });

    const todoListComponent = new TodoListComponent();

    if (resData) {
        for (const data of resData) {
            if (data.completed) {
                todoUlCompleted.insertAdjacentHTML('beforeend', todoListComponent.getCompletedTodoHtml(data));
            } else {
                todoUlUncompleted.insertAdjacentHTML('beforeend', todoListComponent.getUncompletedTodoHtml(data));
            }
        }
    }

    const googleChartHandler = new GoogleChartHandler();
    googleChartHandler.run(resData);
});