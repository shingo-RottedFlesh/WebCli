import { useState, useEffect } from 'react';

// 💡 1. データの型を定義する
interface TaskData {
    text: string; // GoのJSONフィールドに対応
}

// 初期値を、定義した型に合わせる
const initialTask: TaskData = {
    text: "",
};

const Page2 = () => {
    const [task, setTask] = useState<TaskData>(initialTask);
    useEffect(() => {
        fetch('http://localhost:8080/')
            .then(response => response.json())
            .then(data => setTask(data))
            .catch(error =>
                console.error('リクエストエラー:', error));
    }, []);


    return (
        <>
            <div>{task.text}</div>
        </>
    )
}

export default Page2