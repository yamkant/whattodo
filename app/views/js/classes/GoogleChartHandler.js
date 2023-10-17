class GoogleChartHandler {
    getChartDataFormat = (completed_at, content, started_at, ended_at) => {
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

    isValidChartType = (data) => {
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

    run = (resData) => {
        const myDataTable = []
        for (const data of resData) {
            if (!this.isValidChartType(data)) {
                continue;
            }
            myDataTable.push(this.getChartDataFormat(data.completed_at, data.content, data.started_at, data.ended_at)); 
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
}