export interface IContentRes {
	resultObj: {
		containers: IContentContainer[];
	};
}

export interface IContentContainer {
	metadata: { title: string; additionalStreams: IAdditionalStream[] };
	contentId: number;
}

export interface IAdditionalStream {
	hex?: string;
	driverFirstName?: string;
	driverLastName?: string;
	teamName?: string;
	playbackUrl: string;
	title: string | 'DATA' | 'PIT LANE' | 'TRACKER';
}

export interface IHomeRes {
	resultObj: { containers: { retrieveItems: { resultObj: { containers: IHomeContainer[] } } }[] };
}

export interface IHomeContainer {
	id: string;

	metadata: { contentSubtype: string; title: string };
}
