import React from 'react';

import './Header.scss';

function Header(): JSX.Element {
  return (
    <header className="header">
      <div className="header__toolbar">
        <div className="header__left">
          <a href="/" className="brand-link">
            Редактор
          </a>
        </div>
      </div>
    </header>
  );
}

export default Header;
