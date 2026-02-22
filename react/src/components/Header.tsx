import { Link } from 'react-router-dom'

const Header = () => {
    return (
        <header>
            <h1>Header</h1>
            <div>
                <Link to="/home">Home</Link><br />
                <Link to="/page1">Page1</Link><br />
                <Link to="/page2">Page2</Link><br />
                <Link to="/clip/list">ClipList</Link><br />
            </div>
        </header>
    );
};

export default Header;