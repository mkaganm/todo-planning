import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
    const [data, setData] = useState(null);

    useEffect(() => {
        fetch('http://127.0.0.1:3000/api/v1/todo/planning')
            .then((response) => response.json())
            .then((data) => setData(data))
            .catch((error) => console.error("Error fetching data:", error));
    }, []);

    if (!data) return <div>Loading...</div>;

    return (
        <div className="App">
            <h1 className="main-title">Todo Planning</h1>
            {data.data.planning.week_plans.map((weekPlan) => (
                <div key={weekPlan.week_number}>
                    <h2 className="week-title">Week {weekPlan.week_number}</h2>
                    {weekPlan.developer_tasks.map((devTask) => (
                        <div key={devTask.developer.id}>
                            <h3 className="developer-title">{devTask.developer.name}</h3>
                            <table>
                                <thead>
                                <tr>
                                    <th>Task Name</th>
                                    <th>Level</th>
                                    <th>Duration</th>
                                    <th>Estimation</th>
                                </tr>
                                </thead>
                                <tbody>
                                {devTask.tasks.map((task) => (
                                    <tr key={task.name}>
                                        <td>{task.name}</td>
                                        <td>{task.level}</td>
                                        <td>{task.duration}</td>
                                        <td>{task.estimation}</td>
                                    </tr>
                                ))}
                                </tbody>
                            </table>
                        </div>
                    ))}
                </div>
            ))}
        </div>
    );
}

export default App;
