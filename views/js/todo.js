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

const getTodoListHtml = (itemInfo) => {
    function getUuidv4() {
        return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
          var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
            return v.toString(16);
        });
    }
    
    const uuid = getUuidv4();
    const startedAtValue = displayDateTime(itemInfo.started_at);
    const endedAtValue = displayDateTime(itemInfo.ended_at);
    const checkboxLabel = getCheckboxLabelValue(itemInfo.completed);
    const checkedStatus = itemInfo.completed ? "checked" : "";

    return `
        <li data-value=${itemInfo.id} class="todo-item flex items-center text-lg">
            <div class="flex flex-col w-full items-start">
                <div class="flex w-full justify-between">
                    <div class="checkbox-container flex">
                        <input
                            type="checkbox"
                            class="hidden"
                            onclick="onClickTodoItemUpdate(this, 'c')"
                            id="checkbox-${uuid}"
                            ${checkedStatus}
                        >
                        <label for="checkbox-${uuid}">
                            ${checkboxLabel}
                        </label>
                        <div class="text-ellipsis overflow-hidden">${itemInfo.content}</div>
                    </div>
                </div>

                <div class="controller-container w-full flex justify-end space-x-2 text-sm py-2">
                    <div class="start-time-container flex">
                        <button
                            onclick="onClickTodoItemUpdate(this, 's')"
                            class="mb-1 mr-1"
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20" fill="none">
                                <path d="M5 4.98402V19.016C5.00305 19.1907 5.05178 19.3614 5.14135 19.5114C5.23092 19.6613 5.3582 19.7852 5.51052 19.8707C5.66284 19.9561 5.83489 20.0002 6.00955 19.9985C6.1842 19.9968 6.35536 19.9494 6.506 19.861L18.512 12.845C18.6605 12.7595 18.7839 12.6364 18.8696 12.4881C18.9554 12.3397 19.0006 12.1714 19.0006 12C19.0006 11.8287 18.9554 11.6603 18.8696 11.512C18.7839 11.3636 18.6605 11.2405 18.512 11.155L6.506 4.13902C6.35536 4.05062 6.1842 4.00321 6.00955 4.00151C5.83489 3.99982 5.66284 4.0439 5.51052 4.12936C5.3582 4.21483 5.23092 4.3387 5.14135 4.48865C5.05178 4.6386 5.00305 4.80939 5 4.98402Z" fill="#2F2F38"/>
                            </svg>
                        </button>
                        <input class="w-28 border-2 pl-2" value="${startedAtValue}" data-value="${itemInfo.started_at}" type="time">
                    </div>
                    <div class="end-time-container flex">
                        <button
                            onclick="onClickTodoItemUpdate(this, 'e')"
                            class="mb-1 mr-1"
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20" fill="none">
                                <path d="M9 4H8C6.89543 4 6 4.89543 6 6V18C6 19.1046 6.89543 20 8 20H9C10.1046 20 11 19.1046 11 18V6C11 4.89543 10.1046 4 9 4Z" fill="#2F2F38"/>
                                <path d="M16 4H15C13.8954 4 13 4.89543 13 6V18C13 19.1046 13.8954 20 15 20H16C17.1046 20 18 19.1046 18 18V6C18 4.89543 17.1046 4 16 4Z" fill="#2F2F38"/>
                            </svg>
                        </button>
                        <input class="w-28 border-2 pl-2" value="${endedAtValue}" data-value="${itemInfo.ended_at}" type="time">
                    </div>
                    <button onclick="onClickRemoveTodo(this)" class="mb-1">
                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20" fill="none">
                            <path d="M20 6H16V4C16 3.46957 15.7893 2.96086 15.4142 2.58579C15.0391 2.21071 14.5304 2 14 2H10C9.46957 2 8.96086 2.21071 8.58579 2.58579C8.21071 2.96086 8 3.46957 8 4V6H4C3.73478 6 3.48043 6.10536 3.29289 6.29289C3.10536 6.48043 3 6.73478 3 7C3 7.26522 3.10536 7.51957 3.29289 7.70711C3.48043 7.89464 3.73478 8 4 8H5V20C5 20.5304 5.21071 21.0391 5.58579 21.4142C5.96086 21.7893 6.46957 22 7 22H17C17.5304 22 18.0391 21.7893 18.4142 21.4142C18.7893 21.0391 19 20.5304 19 20V8H20C20.2652 8 20.5196 7.89464 20.7071 7.70711C20.8946 7.51957 21 7.26522 21 7C21 6.73478 20.8946 6.48043 20.7071 6.29289C20.5196 6.10536 20.2652 6 20 6ZM10 4H14V6H10V4ZM11 18C11 18.2652 10.8946 18.5196 10.7071 18.7071C10.5196 18.8946 10.2652 19 10 19C9.73478 19 9.48043 18.8946 9.29289 18.7071C9.10536 18.5196 9 18.2652 9 18V10C9 9.73478 9.10536 9.48043 9.29289 9.29289C9.48043 9.10536 9.73478 9 10 9C10.2652 9 10.5196 9.10536 10.7071 9.29289C10.8946 9.48043 11 9.73478 11 10V18ZM15 18C15 18.2652 14.8946 18.5196 14.7071 18.7071C14.5196 18.8946 14.2652 19 14 19C13.7348 19 13.4804 18.8946 13.2929 18.7071C13.1054 18.5196 13 18.2652 13 18V10C13 9.73478 13.1054 9.48043 13.2929 9.29289C13.4804 9.10536 13.7348 9 14 9C14.2652 9 14.5196 9.10536 14.7071 9.29289C14.8946 9.48043 15 9.73478 15 10V18Z" fill="#2F2F38"/>
                        </svg>
                    </button>
                </div>
            </div>
        </li>
    `
}

const toISOStringWithTimezone = date => {
    const tzOffset = -date.getTimezoneOffset();
    const diff = tzOffset >= 0 ? '+' : '-';
    const pad = n => `${Math.floor(Math.abs(n))}`.padStart(2, '0');
    return date.getFullYear() +
        '-' + pad(date.getMonth() + 1) +
        '-' + pad(date.getDate()) +
        'T' + pad(date.getHours()) +
        ':' + pad(date.getMinutes()) +
        ':' + pad(date.getSeconds()) +
        diff + pad(tzOffset / 60) +
        ':' + pad(tzOffset % 60);
};
const onClickTodoItemUpdate = async (target, type) => {
    const todoItemObj = target.closest('.todo-item')
    // start time
    const startInputObj = todoItemObj.querySelector('.start-time-container>input')
    const startTime = toISOStringWithTimezone(new Date())

    // end time
    const endInputObj = todoItemObj.querySelector('.end-time-container>input')
    const endTime = toISOStringWithTimezone(new Date())

    // checked
    const checkboxInputObj = todoItemObj.querySelector('input[type="checkbox"]')

    const data = {}
    switch (type) {
        case 's':
            data.started_at = startTime;
            data.completed = checkboxInputObj.checked;
            break;
        case 'e':
            data.started_at = startInputObj.dataset.value;
            data.ended_at = endTime;
            data.completed = checkboxInputObj.checked;
            break;
        case 'c':
            data.started_at = startInputObj.dataset.value;
            data.ended_at = endInputObj.dataset.value;
            data.completed = checkboxInputObj.checked;
        default:
            break;
    }

    const item_id = target.closest('li').dataset.value;
    const resData = await axios({
        method: 'patch',
        url: `/api/v1/todos/${item_id}/`,
        headers: {
            'Content-Type': 'application/json'
        },
        data: data,
    }).then((response) => {
        return response.data
    }).catch((err) => {
        console.error(err);
    });

    if (resData) {
        switch (type) {
            case 's':
                startInputObj.value = startTime.slice(11, 16);
                break;
            case 'e':
                startInputObj.value = displayDateTime(startInputObj.dataset.value);
                endInputObj.value = endTime.slice(11, 16);
                break;
            case 'c':
                startInputObj.value = displayDateTime(startInputObj.dataset.value);
                endInputObj.value = displayDateTime(endInputObj.dataset.value);
                setCheckBoxByChecked(checkboxInputObj);
                break;
            default:
                break;
        }
    }
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
            return response.data
        })
        .catch((err) => {
            console.error(err);
        });
        if (res) {
            listObj.remove();
        }
    }
}

const setCheckBoxByChecked = (target) => {
    const targetLabel = target.closest('.checkbox-container').querySelector('label');
    targetLabel.innerHTML = getCheckboxLabelValue(target.checked);
}

const onClickAddBtn = async (target) => {
    const todoUlContainer = document.querySelector('.todo-ul-container');
    const todoUl = todoUlContainer.querySelector('.todo-ul');

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
        return response.data
    })
    .catch((err) => {
        console.error(err);
    });

    if (resData) {
        todoUl.insertAdjacentHTML('afterbegin', getTodoListHtml(resData));
        inputObj.value = "";
    }
}

window.addEventListener('load', async () => {
    const todoUlContainer = document.querySelector('.todo-ul-container');
    const todoUl = todoUlContainer.querySelector('.todo-ul');

    const resData = await axios({
        method: 'get',
        url: '/api/v1/todos/',
    })
    .then((response) => {
        return response.data
    })
    .catch((err) => {
        console.error(err);
    });

    if (resData) {
        for (const data of resData) {
            todoUl.insertAdjacentHTML('beforeend', getTodoListHtml(data));
        }
    }
});