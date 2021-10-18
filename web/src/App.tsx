import { useEffect, useState } from 'react';
import styles from './App.module.scss';
import { getEvent, getPlaylistURL } from './core/http/f1tv.http';
import { AppEntry } from './core/types/app.types';
import { mapContentToApp } from './core/utils/content.mapper';

export const App: React.FC = () => {
	const [entries, setEntries] = useState<AppEntry[]>([]);
	useEffect(() => {
		const urlSearchParams = Object.fromEntries(
			new URLSearchParams(window.location.search).entries()
		);
		getEvent(urlSearchParams.contentId ?? '1000003972')
			.then((res) => mapContentToApp(res))
			.then((content) =>
				setEntries(content.sort((a, b) => b.hex.localeCompare(a.hex)))
			);
	}, []);

	const onClick = (contentId: string, channelId?: string) =>
		getPlaylistURL(contentId, channelId).then(({ url }) =>
			navigator.clipboard.writeText(url)
		);

	return (
		<div className={styles.App}>
			{entries.map(({ title, teamName, hex, contentId, channelId }, i) => (
				<span
					key={i}
					style={{ color: hex }}
					onClick={() => onClick(contentId, channelId)}
					className={styles.item}
				>
					{`${teamName ? `${teamName} - ` : ''}${title}`}
				</span>
			))}
		</div>
	);
};
