import React from 'react';
import { Navigate, useLocation } from 'react-router-dom';

/* 
 *  このコンポーネントは、ログインしているかどうかを判定し、
 *  ログインしていない場合はログインページにリダイレクトする
 *  ログインしている場合は、children要素を表示する
 */
// 引数で渡されるprops の children要素を受け取る
function RequireAuth({ children }: { children: React.ReactNode }) {
    const isAuthenticated = localStorage.getItem('isAuthenticated') === 'true';
    const location = useLocation();

    if (!isAuthenticated) {
        // ログインしていない場合、ログインページにリダイレクト
        // state={{ from: location }} で、ログイン後に元のページに戻れるように情報を渡す（今回は未使用だが一般的）
        return <Navigate to="/" state={{ from: location }} replace />;
    }

    return children;
}

export default RequireAuth;