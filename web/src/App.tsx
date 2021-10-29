import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { HomeView } from './views/Home.view';
import { PlaylistView } from './views/Playlist.view';

export const App: React.FC = () => (
	<Router>
		<Switch>
			<Route path="/events/:eventId">
				<PlaylistView />
			</Route>
			<Route path="/">
				<HomeView />
			</Route>
		</Switch>
	</Router>
);
