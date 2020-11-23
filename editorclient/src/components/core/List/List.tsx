import React, { forwardRef } from 'react';
import classNames from 'classnames';

import './List.css';

interface ListProps extends React.HTMLAttributes<HTMLUListElement> {
  className?: string;
}

const List = forwardRef(function List(
  { className, ...rest }: ListProps,
  ref: React.Ref<HTMLUListElement>
) {
  return <ul ref={ref} className={classNames('list', className)} {...rest} />;
});

export default List;
