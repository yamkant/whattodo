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

const getChartDataFormat = (completed_at, content, started_at, ended_at) => {
    const completedDate = new Date(completed_at);
    const startDate = new Date(started_at);
    const endDate = new Date(ended_at);
    return [
        completedDate.getFullYear() + '-' + (completedDate.getMonth() + 1) + '-' + completedDate.getDate(),
        content,
        new Date(0,0,0,startDate.getHours(),startDate.getMinutes(),0),
        new Date(0,0,0,endDate.getHours(),endDate.getMinutes(),0),
    ]
}

const isValidChartType = (data) => {
    if (!data.completed) {
        return false;
    }
    if (data.started_at === '0001-01-01T00:00:00Z') {
        return false;
    }
    if (data.ended_at === '0001-01-01T00:00:00Z') {
        return false;
    }
    if (data.started_at >= data.ended_at) {
        return false;
    }
    return true;
}

const setChart = (resData) => {
    const myDataTable = []
    for (const data of resData) {
        if (!isValidChartType(data)) {
            continue;
        }
        myDataTable.push(getChartDataFormat(data.completed_at, data.content, data.started_at, data.ended_at)); 
    }

    google.charts.load("current", {packages:["timeline"]});
    google.charts.setOnLoadCallback(getHandleredDataTable);
    function getHandleredDataTable() {
        const container = document.getElementById('todo-chart');
        const chart = new google.visualization.Timeline(container);
        const options = {
            timeline: { colorByRowLabel: true },
            alternatingRowStyle: false
        };
        const dataTable = new google.visualization.DataTable();
        dataTable.addColumn({ type: 'string', id: 'Room' });
        dataTable.addColumn({ type: 'string', id: 'Name' });
        dataTable.addColumn({ type: 'date', id: 'Start' });
        dataTable.addColumn({ type: 'date', id: 'End' });

        dataTable.addRows(myDataTable);

        chart.draw(dataTable, options);
    }

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

    setChart(resData);
});