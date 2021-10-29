import { AppEvent } from '../types/app.types';
import { IHomeRes } from '../types/f1tv.types';

export const mapHomeToApp = (homeRes: IHomeRes): AppEvent[] =>
	homeRes.resultObj.containers
		.flatMap(({ retrieveItems }) =>
			retrieveItems.resultObj.containers.filter(({ metadata }) => metadata?.contentSubtype === 'LIVE')
		)
		.map(({ id, metadata }) => ({ title: metadata?.title, eventId: id }));
