import { FC } from 'react';
import { AppEntry } from '../../../core/types/app.types';
import { PlaylistEntry } from '../../atoms/playlist-entry/PlaylistEntry.component';
import styles from './PlaylistScreen.module.scss';

export interface PlaylistScreenProps {
	entries: AppEntry[];
	onLoadStream: (contentId: string, channelId?: string) => void;
}

export const PlaylistScreen: FC<PlaylistScreenProps> = ({
	entries,
	onLoadStream,
}) => (
	<div className={styles.root}>
		{entries.map(({ title, hex, contentId, channelId }, i) => (
			<PlaylistEntry
				key={i}
				color={hex}
				text={title}
				onClick={() => onLoadStream(contentId, channelId)}
			/>
		))}
	</div>
);
