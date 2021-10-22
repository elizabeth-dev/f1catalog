import { BrowserRouter as Router, Route } from 'react-router-dom';
import { PlaylistView } from './views/Playlist.view';

export const App: React.FC = () => (
	<Router>
		<Route path="/events/:eventId">
			<PlaylistView />
		</Route>
	</Router>
);
