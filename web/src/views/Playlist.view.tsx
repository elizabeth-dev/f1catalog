import { FC, useEffect, useState } from 'react';
import { useParams } from 'react-router';
import { PlaylistScreen } from '../components/screens/playlist/PlaylistScreen.component';
import { getEvent } from '../core/http/f1tv.http';
import { AppEntry } from '../core/types/app.types';
import { mapContentToApp } from '../core/utils/content.mapper';
import { genPlaybackUrl } from '../core/utils/playbackUrl.util';

export const PlaylistView: FC = () => {
	const { eventId } = useParams<{ eventId: string }>();

	const [entries, setEntries] = useState<AppEntry[]>([]);

	useEffect(() => {
		getEvent(eventId)
			.then((res) => mapContentToApp(res))
			.then((content) => setEntries(content.sort((a, b) => b.hex.localeCompare(a.hex))));
	}, [eventId]);

	const onLoadStream = (contentId: string, channelId?: string) =>
		navigator.clipboard.writeText(genPlaybackUrl(contentId, channelId));

	return <PlaylistScreen entries={entries} onLoadStream={onLoadStream} />;
};
