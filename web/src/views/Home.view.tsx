import { FC, useCallback, useEffect, useReducer } from 'react';
import { HomeScreen } from '../components/screens/home/HomeScreen.component';
import { getPage, PAGE_HOME } from '../core/http/f1tv.http';
import { AppEvent } from '../core/types/app.types';
import { mapHomeToApp } from '../core/utils/page.mapper';

interface HomeViewState {
	loading: boolean;
	events: AppEvent[];
}

interface HomeViewAction {
	type: string;
	payload?: { events: AppEvent[] };
}

const initialState = { loading: true, events: [] };
const reducer = (state: HomeViewState, action: HomeViewAction) => {
	switch (action.type) {
		case 'RELOADED':
			return { ...state, loading: false, events: action.payload?.events ?? [] };
		case 'RELOADING':
			return { ...state, loading: true };
		default:
			return state;
	}
};

// const urlSearchParams = Object.fromEntries(new URLSearchParams(window.location.search).entries());

export const HomeView: FC = () => {
	const [{ events, loading }, dispatch] = useReducer(reducer, initialState);

	const reload = useCallback(async () => {
		dispatch({ type: 'RELOADING' });

		const res = await getPage(PAGE_HOME);
		const events = mapHomeToApp(res);

		dispatch({ type: 'RELOADED', payload: { events } });
		// if (!urlSearchParams.noAutoReload && events.length === 0) setTimeout(reload, 120000);
	}, []);

	useEffect(() => {
		reload();
	}, [reload]);

	return <HomeScreen events={events} loading={loading} />;
};
