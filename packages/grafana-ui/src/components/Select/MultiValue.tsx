import * as React from 'react';

import { t } from '@grafana/i18n';

import { useTheme2 } from '../../themes/ThemeContext';
import { IconButton, Props as IconButtonProps } from '../IconButton/IconButton';

import { getSelectStyles } from './getSelectStyles';

interface MultiValueContainerProps {
  innerProps: JSX.IntrinsicElements['div'];
}

export const MultiValueContainer = ({ innerProps, children }: React.PropsWithChildren<MultiValueContainerProps>) => {
  const theme = useTheme2();
  const styles = getSelectStyles(theme);

  return (
    <div {...innerProps} className={styles.multiValueContainer}>
      {children}
    </div>
  );
};

export type MultiValueRemoveProps = {
  innerProps: IconButtonProps;
};

export const MultiValueRemove = ({ children, innerProps }: React.PropsWithChildren<MultiValueRemoveProps>) => {
  const theme = useTheme2();
  const styles = getSelectStyles(theme);
  return (
    <IconButton
      {...innerProps}
      name="times"
      size="sm"
      className={styles.multiValueRemove}
      tooltip={t('grafana-ui.select.multi-value-remove', 'Remove')}
    />
  );
};
