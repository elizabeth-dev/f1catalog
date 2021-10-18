export interface IContentRes {
	resultObj: {
		containers: IContainer[];
	};
}

export interface IContainer {
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
