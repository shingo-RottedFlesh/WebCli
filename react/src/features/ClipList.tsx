import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import '../App.css'

// jsonの受け取り型を定義
interface ClipList {
    clipId: string;
    clipName: string;
}

// 初期値設定
const initialClipList: ClipList[] = [];

function Home() {
    const [clipList, setClipList] = useState<ClipList[]>(initialClipList);
    useEffect(() => {
        fetch('http://localhost:8080/clip/list', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        })
            .then(response => response.json())
            .then(data => setClipList(data))
            .catch(error =>
                console.error('リクエストエラー:', error));
    }, []);

    return (
        <>

            <div>
                <Link to="/home">Home</Link><br />
            </div>
            <h1>クリップ一覧</h1>
            <div>
                {clipList.map((clip) => (
                    <div key={clip.clipId}>
                        <p>クリップID: {clip.clipId}</p>
                        <p>クリップ名: {clip.clipName}</p>
                    </div>
                ))}
            </div>

        </>
    )
}

export default Home