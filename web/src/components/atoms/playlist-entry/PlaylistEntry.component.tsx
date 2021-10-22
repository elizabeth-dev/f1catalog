import { FC } from 'react';
import styles from './PlaylistEntry.module.scss';

export interface PlaylistEntryProps {
	text: string;
	color: string;
	onClick: React.MouseEventHandler<HTMLSpanElement>;
}
export const PlaylistEntry: FC<PlaylistEntryProps> = ({
	text,
	color,
	onClick,
}) => (
	<span style={{ color }} onClick={onClick} className={styles.root}>
		{text}
	</span>
);
