export interface AppEntry {
	title: string;
	hex: string;
	contentId: string;
	channelId?: string;
	driver: boolean;
}

export interface AppEvent {
	title: string;
	eventId: string;
}
