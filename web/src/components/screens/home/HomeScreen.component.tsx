import { FC } from 'react';
import { Link } from 'react-router-dom';
import styles from './HomeScreen.module.scss';

export interface HomeScreenProps {
	events: { title: string; eventId: string }[];
	loading: boolean;
}

export const HomeScreen: FC<HomeScreenProps> = ({ events, loading }) => (
	<div className={styles.root}>
		{events.map(({ title, eventId }, i) => (
			<Link key={i} to={`/events/${eventId}`} className={styles.link}>
				{title}
			</Link>
		))}
		{!loading && events.length === 0 && <span className={styles.disclaimer}>No live events available</span>}
		{loading && <span className={styles.disclaimer}>Loading...</span>}
	</div>
);
