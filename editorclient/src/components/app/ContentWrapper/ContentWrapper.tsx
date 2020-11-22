import React from 'react';

import './ContentWrapper.scss';

interface Props {
  children: JSX.Element;
}

function ContentWrapper({ children }: Props): JSX.Element {
  return <main className="content-wrapper">{children}</main>;
}

export default ContentWrapper;
