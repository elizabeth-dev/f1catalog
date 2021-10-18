import { AppEntry } from '../types/app.types';
import { IContentRes } from '../types/f1tv.types';
const DEFAULT_COLOR = '#ffffff';
export const mapContentToApp = ({ resultObj }: IContentRes): AppEntry[] => [
	{
		hex: DEFAULT_COLOR,
		title: 'World Feed',
		contentId: resultObj.containers[0].contentId.toString(),
	},
	...resultObj.containers[0].metadata.additionalStreams.map(
		({
			hex,
			playbackUrl,
			title,
			driverFirstName,
			driverLastName,
			teamName,
		}) => ({
			hex: hex ?? DEFAULT_COLOR,
			title:
				driverFirstName && driverLastName
					? `${driverFirstName} ${driverLastName}`
					: title,
			contentId: [...playbackUrl.matchAll(/contentId=(\d+)/gi)][0][1],
			channelId: [...playbackUrl.matchAll(/channelId=(\d+)/gi)][0][1],
			teamName,
		})
	),
];
