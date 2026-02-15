import React from 'react';
import { Navigate, useLocation } from 'react-router-dom';

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