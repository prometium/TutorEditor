import React, { forwardRef } from 'react';
import classNames from 'classnames';

import './ListItem.css';

interface Props extends React.HTMLAttributes<HTMLLIElement> {
  className?: string;
}

const ListItem = forwardRef(function ListItem(
  { className, ...rest }: Props,
  ref: React.Ref<HTMLLIElement>
) {
  return (
    <li ref={ref} className={classNames('list-item', className)} {...rest} />
  );
});

export default ListItem;
