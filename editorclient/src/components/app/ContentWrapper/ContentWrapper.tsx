import React from 'react';

import './ContentWrapper.scss';

type ContentWrapperProps = {
  children: React.ReactNode;
};

function ContentWrapper({ children }: ContentWrapperProps): JSX.Element {
  return <main className="content-wrapper">{children}</main>;
}

export default ContentWrapper;
