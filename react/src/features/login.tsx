import { useState, type FormEvent } from "react"
import { useNavigate } from 'react-router-dom';

const Login = () => {
    // ログインの入力項目の変数設定
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const [error, setError] = useState('');

    // バリデーション部分の要素操作用
    // const validationRef = useRef<HTMLDivElement>(null);

    const navigate = useNavigate();

    const submitForm = (event: FormEvent) => {
        // ページのリロードを防ぐ
        event.preventDefault();

        // エラーをリセット
        setError('');

        fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password })
        })
            .then(response => {
                if (!response.ok) {
                    // サーバーエラー等の場合もここでキャッチしてJSONとしてパースを試みるか、エラーを投げる
                    // 今回はmain.goで200 OKまたは500を返しているので、そのままparseへ
                    return response.json().catch(() => {
                        throw new Error('サーバー通信エラーが発生しました');
                    });
                }
                return response.json();
            })
            .then(data => {
                console.log('Response data:', data);
                if (data.success) {
                    // ログイン成功
                    navigate('/home');
                } else {
                    // ログイン失敗（バリデーションエラーやシステムエラー）
                    setError(data.message || "ログインに失敗しました");
                }
            })
            .catch(error => {
                console.error('リクエストエラー:', error);
                setError("通信エラーが発生しました");
            });
    }

    return (
        <>
            <div>ログイン画面</div>
            <form onSubmit={submitForm}>
                <div>
                    <div id='validation'
                        style={{ color: '#ff0000' }}>{error}</div>
                    <label htmlFor="username">ユーザ名</label>
                    <input type="text" name="username" id="username" value={username}
                        onChange={(e) => setUsername(e.target.value)} required />
                    <br />
                    <label htmlFor="password">パスワード</label>
                    <input type="password" id="password" value={password}
                        onChange={(e) => setPassword(e.target.value)} required />
                </div>
                <input type="submit" />
            </form>
        </>
    );
}


export default Login