class TodoListComponent {
    getUuidv4() {
        return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
            return v.toString(16);
        });
    }
        
    getListContentHtml = (checkedStatus, checkboxLabel, content) => {
        const uuid = this.getUuidv4();
        return `
            <div class="flex flex-col w-full items-start text-sm md:text-lg">
                <div class="flex w-full justify-between">
                    <div class="checkbox-container flex">
                        <input
                            type="checkbox"
                            class="hidden"
                            onclick="TodoAPIHandler.updateTodoCompleted(this); location.reload();"
                            id="checkbox-${uuid}"
                            ${checkedStatus}
                        >
                        <label for="checkbox-${uuid}">
                            ${checkboxLabel}
                        </label>
                        <input
                            class="todo-content w-72 md:w-80 text-ellipsis overflow-hidden cursor-pointer"
                            onchange="TodoAPIHandler.updateTodoContent(this);"
                            value='${content}'
                        >
                    </div>
                </div>
            </div>
        `
    }

    getListTimeBoxHtml = (startedAt, endedAt) => {
        const startedAtValue = displayDateTime(startedAt);
        const endedAtValue = displayDateTime(endedAt);
        return `
            <div class="controller-container w-full flex justify-end space-x-2 text-sm py-2">
                <div class="start-time-container flex">
                    <button
                        onclick="TodoAPIHandler.updateTodoStartAt(this)"
                        class="mb-1 mr-1"
                    >
                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20" fill="none">
                            <path d="M5 4.98402V19.016C5.00305 19.1907 5.05178 19.3614 5.14135 19.5114C5.23092 19.6613 5.3582 19.7852 5.51052 19.8707C5.66284 19.9561 5.83489 20.0002 6.00955 19.9985C6.1842 19.9968 6.35536 19.9494 6.506 19.861L18.512 12.845C18.6605 12.7595 18.7839 12.6364 18.8696 12.4881C18.9554 12.3397 19.0006 12.1714 19.0006 12C19.0006 11.8287 18.9554 11.6603 18.8696 11.512C18.7839 11.3636 18.6605 11.2405 18.512 11.155L6.506 4.13902C6.35536 4.05062 6.1842 4.00321 6.00955 4.00151C5.83489 3.99982 5.66284 4.0439 5.51052 4.12936C5.3582 4.21483 5.23092 4.3387 5.14135 4.48865C5.05178 4.6386 5.00305 4.80939 5 4.98402Z" fill="#2F2F38"/>
                        </svg>
                    </button>
                    <input class="w-28 border-2 pl-2" value="${startedAtValue}" data-value="${startedAt}" type="time" disabled>
                </div>
                <div class="end-time-container flex">
                    <button
                        onclick="TodoAPIHandler.updateTodoEndAt(this)"
                        class="mb-1 mr-1"
                    >
                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20" fill="none">
                            <path d="M9 4H8C6.89543 4 6 4.89543 6 6V18C6 19.1046 6.89543 20 8 20H9C10.1046 20 11 19.1046 11 18V6C11 4.89543 10.1046 4 9 4Z" fill="#2F2F38"/>
                            <path d="M16 4H15C13.8954 4 13 4.89543 13 6V18C13 19.1046 13.8954 20 15 20H16C17.1046 20 18 19.1046 18 18V6C18 4.89543 17.1046 4 16 4Z" fill="#2F2F38"/>
                        </svg>
                    </button>
                    <input class="w-28 border-2 pl-2" value="${endedAtValue}" data-value="${endedAt}" type="time" disabled>
                </div>
                <button onclick="onClickRemoveTodo(this)" class="mb-1">
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20" fill="none">
                        <path d="M20 6H16V4C16 3.46957 15.7893 2.96086 15.4142 2.58579C15.0391 2.21071 14.5304 2 14 2H10C9.46957 2 8.96086 2.21071 8.58579 2.58579C8.21071 2.96086 8 3.46957 8 4V6H4C3.73478 6 3.48043 6.10536 3.29289 6.29289C3.10536 6.48043 3 6.73478 3 7C3 7.26522 3.10536 7.51957 3.29289 7.70711C3.48043 7.89464 3.73478 8 4 8H5V20C5 20.5304 5.21071 21.0391 5.58579 21.4142C5.96086 21.7893 6.46957 22 7 22H17C17.5304 22 18.0391 21.7893 18.4142 21.4142C18.7893 21.0391 19 20.5304 19 20V8H20C20.2652 8 20.5196 7.89464 20.7071 7.70711C20.8946 7.51957 21 7.26522 21 7C21 6.73478 20.8946 6.48043 20.7071 6.29289C20.5196 6.10536 20.2652 6 20 6ZM10 4H14V6H10V4ZM11 18C11 18.2652 10.8946 18.5196 10.7071 18.7071C10.5196 18.8946 10.2652 19 10 19C9.73478 19 9.48043 18.8946 9.29289 18.7071C9.10536 18.5196 9 18.2652 9 18V10C9 9.73478 9.10536 9.48043 9.29289 9.29289C9.48043 9.10536 9.73478 9 10 9C10.2652 9 10.5196 9.10536 10.7071 9.29289C10.8946 9.48043 11 9.73478 11 10V18ZM15 18C15 18.2652 14.8946 18.5196 14.7071 18.7071C14.5196 18.8946 14.2652 19 14 19C13.7348 19 13.4804 18.8946 13.2929 18.7071C13.1054 18.5196 13 18.2652 13 18V10C13 9.73478 13.1054 9.48043 13.2929 9.29289C13.4804 9.10536 13.7348 9 14 9C14.2652 9 14.5196 9.10536 14.7071 9.29289C14.8946 9.48043 15 9.73478 15 10V18Z" fill="#2F2F38"/>
                    </svg>
                </button>
            </div>
        `
    }

    getListHistoryBoxHtml = (completedAt, startedAt, endedAt) => {
        function getYearMonthDayDate() {
            const tmpDate = completedAt.split('T')[0].split('-');
            return tmpDate.join('.')
        }

        function getTimeDisplayString(diffMsec) {
            const diffHour = Math.floor(diffMsec / (60 * 60 * 1000));
            const diffHourDisp = diffHour > 0 ? `${diffHour}시간 ` : '';
            const diffMin = Math.floor(diffMsec / (60 * 1000) % 60);
            const diffMinDisp = diffMin > 0 ? `${diffMin}분` : '';
            if (diffHour + diffMin !== 0 && (diffHour >= 0 || diffMin >= 0)) {
                return `소요시간: ${diffHourDisp}${diffMinDisp}`
            } else {
                return ""
            }
        }

        const diffMsec = new Date(endedAt) - new Date(startedAt);
        return `
            <div class="controller-container w-full flex justify-between space-x-2 text-sm py-2">
                <div class="flex flex-row">
                    <div>${getYearMonthDayDate(completedAt)}</div>
                    <div class="ml-5">${getTimeDisplayString(diffMsec)}</div>
                </div>
                <button onclick="onClickRemoveTodo(this)" class="mb-1">
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20" fill="none">
                        <path d="M20 6H16V4C16 3.46957 15.7893 2.96086 15.4142 2.58579C15.0391 2.21071 14.5304 2 14 2H10C9.46957 2 8.96086 2.21071 8.58579 2.58579C8.21071 2.96086 8 3.46957 8 4V6H4C3.73478 6 3.48043 6.10536 3.29289 6.29289C3.10536 6.48043 3 6.73478 3 7C3 7.26522 3.10536 7.51957 3.29289 7.70711C3.48043 7.89464 3.73478 8 4 8H5V20C5 20.5304 5.21071 21.0391 5.58579 21.4142C5.96086 21.7893 6.46957 22 7 22H17C17.5304 22 18.0391 21.7893 18.4142 21.4142C18.7893 21.0391 19 20.5304 19 20V8H20C20.2652 8 20.5196 7.89464 20.7071 7.70711C20.8946 7.51957 21 7.26522 21 7C21 6.73478 20.8946 6.48043 20.7071 6.29289C20.5196 6.10536 20.2652 6 20 6ZM10 4H14V6H10V4ZM11 18C11 18.2652 10.8946 18.5196 10.7071 18.7071C10.5196 18.8946 10.2652 19 10 19C9.73478 19 9.48043 18.8946 9.29289 18.7071C9.10536 18.5196 9 18.2652 9 18V10C9 9.73478 9.10536 9.48043 9.29289 9.29289C9.48043 9.10536 9.73478 9 10 9C10.2652 9 10.5196 9.10536 10.7071 9.29289C10.8946 9.48043 11 9.73478 11 10V18ZM15 18C15 18.2652 14.8946 18.5196 14.7071 18.7071C14.5196 18.8946 14.2652 19 14 19C13.7348 19 13.4804 18.8946 13.2929 18.7071C13.1054 18.5196 13 18.2652 13 18V10C13 9.73478 13.1054 9.48043 13.2929 9.29289C13.4804 9.10536 13.7348 9 14 9C14.2652 9 14.5196 9.10536 14.7071 9.29289C14.8946 9.48043 15 9.73478 15 10V18Z" fill="#2F2F38"/>
                    </svg>
                </button>
            </div>
        `
    }

    getUncompletedTodoHtml = (itemInfo) => {
        const checkboxLabel = getCheckboxLabelValue(itemInfo.completed);
        const checkedStatus = itemInfo.completed ? "checked" : "";

        return `
            <li data-value=${itemInfo.id} class="todo-item flex items-center text-lg">
                <div class="flex flex-col w-full items-start">
                    ${this.getListContentHtml(checkedStatus, checkboxLabel, itemInfo.content)}
                    ${this.getListTimeBoxHtml(itemInfo.started_at, itemInfo.ended_at)}
                </div>
            </li>
        `
    }

    getCompletedTodoHtml = (itemInfo) => {
        const checkboxLabel = getCheckboxLabelValue(itemInfo.completed);
        const checkedStatus = itemInfo.completed ? "checked" : "";

        return `
            <li data-value=${itemInfo.id} class="todo-item flex items-center text-lg">
                <div class="flex flex-col w-full items-start">
                    ${this.getListContentHtml(checkedStatus, checkboxLabel, itemInfo.content)}
                    ${this.getListHistoryBoxHtml(itemInfo.completed_at, itemInfo.started_at, itemInfo.ended_at)}
                </div>
            </li>
        `
    }
}
