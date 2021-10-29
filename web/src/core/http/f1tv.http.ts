import { IContentRes, IHomeRes } from '../types/f1tv.types';

const CORS_PROXY_URL = 'https://us-central1-f1tv-325811.cloudfunctions.net/CorsProxy';
const PLAYLIST_URL = 'https://us-central1-f1tv-325811.cloudfunctions.net/GetPlaylistURL';

export const PAGE_HOME = '395';

// UNUSED
export const getPlaybackURL = (contentId: string, channelId?: string): Promise<{ url: string }> =>
	fetch(`${PLAYLIST_URL}?contentId=${contentId}${channelId ? `&channelId=${channelId}` : ''}`).then((res) =>
		res.json()
	);

export const getEvent = (contentId: string): Promise<IContentRes> =>
	fetch(`${CORS_PROXY_URL}?uri=/2.0/R/ENG/BIG_SCREEN_HLS/ALL/CONTENT/VIDEO/${contentId}/F1_TV_Pro_Annual/2`).then(
		(res) => res.json()
	);

export const getPage = (pageId: string): Promise<IHomeRes> =>
	fetch(`${CORS_PROXY_URL}?uri=/2.0/R/ENG/BIG_SCREEN_HLS/ALL/PAGE/${pageId}/F1_TV_Pro_Annual/2`).then((res) =>
		res.json()
	);
