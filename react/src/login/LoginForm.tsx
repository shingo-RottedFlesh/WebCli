import {useState, useRef, type FormEvent} from "react"
import { useNavigate } from 'react-router-dom';

const LoginForm = () => {
    // ログインの入力項目の変数設定
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    // バリデーション部分の要素操作用
    const validationRef = useRef<HTMLDivElement>(null);

    const navigate = useNavigate();

    const submitForm = (event: FormEvent) => {
        // ページのリロードを防ぐ
        event.preventDefault();
        
        // Todo：Golang側のログイン認証ロジックを使う

        if (username === 'shingo' && password === 'password' ){
            navigate('/home')
        } else {
            if (validationRef.current) {
                validationRef.current.innerText = 'ログインに失敗しました';
            }
        }
    }


    return(
        <>
        <div>ログイン画面</div>
        <form onSubmit={submitForm}>
            <div>
                <div id='validation' 
                    ref={validationRef}
                    style={{ color: '#ff0000'}}></div>
                <label htmlFor="username">ユーザ名</label>
                <input type="text" id="username" value={username} 
                    onChange={(e) => setUsername(e.target.value)} required />
                <br/>
                <label htmlFor="password">パスワード</label>
                <input type="password" id="password" value={password} 
                    onChange={(e) => setPassword(e.target.value)} required />
            </div>
            <input type="submit" />
        </form>
        </>
    );
}


export default LoginForm