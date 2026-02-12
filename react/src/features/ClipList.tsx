import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import '../App.css'

// jsonの受け取り型を定義
interface FileList {
    fileId: string;
    fileName: string;
}

// 初期値設定
const initialFileList: FileList[] = [];

function Home() {
    const [fileList, setFileList] = useState<FileList[]>(initialFileList);
    useEffect(() => {
        fetch('http://localhost:8080/clip/list', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        })
            .then(response => response.json())
            .then(data => setFileList(data))
            .catch(error =>
                console.error('リクエストエラー:', error));
    }, []);

    return (
        <>

            <div>
                <Link to="/home">Home</Link><br />
            </div>
            <h1>ファイル一覧</h1>
            <div>
                {fileList.map((file) => (
                    <div key={file.fileId}>
                        <p>ファイルID: {file.fileId}</p>
                        <p>ファイル名: {file.fileName}</p>
                    </div>
                ))}
            </div>

        </>
    )
}

export default Home