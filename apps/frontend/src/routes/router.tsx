import { createBrowserRouter } from 'react-router-dom';
import { LoginRoot, RegisterRoot } from '../features/auth/auth-routes';
import { JobsRoot, NewJobRoot } from '../features/jobs/jobs-routes';
import { LandingRoute } from '../features/landing/landing-routes';
import { AppLayout } from '../features/layout/AppLayout';
import { View } from '../features/plugs/View';

export const router = createBrowserRouter([
	{
		path: '/',
		element: <LandingRoute />,
	},
	{
		path: '/login',
		element: <LoginRoot />,
	},
	{
		path: '/register',
		element: <RegisterRoot />,
	},
	{
		path: '/app',
		element: <AppLayout />,
		children: [
			{
				path: '/app',
			},
			{
				path: 'plugs',
				element: <View />,
			},
			{
				path: 'jobber',
				element: <JobsRoot />,
			},
			{
				path: 'jobber/new',
				element: <NewJobRoot />,
			},
			{
				path: 'chat',
				element: <div>Hello</div>,
			},
		],
	},
]);
