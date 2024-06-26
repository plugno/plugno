'use client';

import { Provider } from 'react-redux';
import { store } from '@store/store';
import './globals.css';

export default function RootLayout({ children }: { children: React.ReactNode }) {
	return (
		<html lang="en">
			<head />
			<body>
				<Provider store={store}>
					<div>{children}</div>
				</Provider>
			</body>
		</html>
	);
}
