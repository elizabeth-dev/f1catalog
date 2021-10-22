const PLAYBACK_BASE = 'https://us-central1-f1tv-325811.cloudfunctions.net/FetchPlayback';

export const genPlaybackUrl = (contentId: string, channelId?: string) =>
	`${PLAYBACK_BASE}?contentId=${contentId}${channelId ? `&channelId=${channelId}` : ''}`;
