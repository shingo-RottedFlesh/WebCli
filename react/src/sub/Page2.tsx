import { useState, useEffect } from 'react';


const Page2 = () => {
    const [task, setTask] = useState("");
    useEffect(() => {
fetch('http://localhost:8080')
  .then(response => response.json())
  .then(data => setTask(data))
  .catch(error => 
    console.error('リクエストエラー:', error));
    }, []);


    return (
        <>
            <div>{task.Text}</div>
        </>
        )
    }

export default Page2