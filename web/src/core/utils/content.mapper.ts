import { AppEntry } from '../types/app.types';
import { IContentRes } from '../types/f1tv.types';

const DEFAULT_COLOR = '#ffffff';

const TITLE_MAPPING: { [key: string]: string } = {
	'PIT LANE': 'Pit Lane Channel',
	TRACKER: 'Driver Tracker',
	DATA: 'Data Channel',
};

const generateTitle = (title: string, teamName?: string, driverFirstName?: string, driverLastName?: string): string =>
	`${teamName ? `${teamName} - ` : ''}${
		driverFirstName && driverLastName ? `${driverFirstName} ${driverLastName}` : TITLE_MAPPING[title]
	}`;

export const mapContentToApp = ({ resultObj }: IContentRes): AppEntry[] => [
	{
		hex: DEFAULT_COLOR,
		title: 'World Feed',
		contentId: resultObj.containers[0].contentId.toString(),
		driver: false,
	},
	...(resultObj.containers[0].metadata.additionalStreams?.map(
		({ hex, playbackUrl, title, driverFirstName, driverLastName, teamName }) => ({
			hex: hex ?? DEFAULT_COLOR,
			title: generateTitle(title, teamName, driverFirstName, driverLastName),
			contentId: [...playbackUrl.matchAll(/contentId=(\d+)/gi)][0][1],
			channelId: [...playbackUrl.matchAll(/channelId=(\d+)/gi)][0][1],
			teamName,
			driver: !!driverFirstName,
		})
	) ?? []),
];
