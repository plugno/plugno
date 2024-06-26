'use client';

import { UserResponse } from '@typings/user';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { useState } from 'react';

export default function LoginPage() {
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');
	const router = useRouter();

	const handleLogin = async () => {
		fetch('http://localhost:6001/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			credentials: 'include',
			body: JSON.stringify({
				email,
				password,
			}),
		})
			.then((res) => res.json())
			.then((data: UserResponse) => {
				if (data.isSuccess) router.push('/app/plugs');
			});
	};

	const redirect_uri = 'https://5d8c-92-220-236-200.eu.ngrok.io/login';

	const handleFacebookLogin = async () => {
		fetch(
			`https://www.facebook.com/v15.0/dialog/oauth?client_id=2272911636269824&redirect_uri=${redirect_uri}&state="{st=state123abc,ds=123456789}"`,
		)
			.then((res) => res.json())
			.then((data) => {
				console.log('fb data:', data);
			});
	};

	return (
		<div className="bg-neutral-900 min-h-screen flex flex-col justify-center py-12 sm:px-6 lg:px-8">
			<div className="sm:mx-auto sm:w-full sm:max-w-md">
				<div>
					<h2 className="text-center text-3xl text-white font-bold tracking-tight">Logg inn</h2>
					<p className="mt-2 text-center text-sm text-gray-300">
						Eller{' '}
						<Link href="/register" className="font-medium text-rose-500 hover:text-red-600">
							lag en bruker.
						</Link>
					</p>
				</div>
			</div>
			<div className="mt-8 shadow-sm sm:mx-auto sm:w-full sm:max-w-md">
				<div className="bg-neutral-800 py-8 px-4 shadow sm:rounded-lg sm:px-10 space-y-4">
					<div>
						<label htmlFor="email" className="block text-sm font-medium text-gray-300">
							E-post
						</label>
						<div className="mt-1">
							<input
								id="email"
								name="email"
								type="email"
								autoComplete="email"
								required
								className="block w-full outline-none text-white rounded-md bg-neutral-700 border border-neutral-600 px-3 py-2 placeholder-gray-400 shadow-sm sm:text-sm"
								value={email}
								onChange={(e) => setEmail(e.currentTarget.value)}
							/>
						</div>
					</div>
					<div>
						<label htmlFor="password" className="block text-sm font-medium text-gray-300">
							Passord
						</label>
						<div className="mt-1">
							<input
								id="password"
								name="password"
								type="password"
								autoComplete="current-password"
								required
								className="block w-full outline-none text-white rounded-md bg-neutral-700 border border-neutral-600 px-3 py-2 placeholder-gray-400 shadow-sm sm:text-sm"
								value={password}
								onChange={(e) => setPassword(e.currentTarget.value)}
							/>
						</div>
					</div>
					<div className="space-y-8">
						<button
							onClick={handleLogin}
							className="flex w-full justify-center rounded-md border border-rose-600/60 bg-rose-500 bg-opacity-60 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-rose-600/60 focus:outline-none focus:ring-2 focus:ring-rose-500 focus:ring-offset-2"
						>
							Logg inn
						</button>

						<button
							onClick={handleFacebookLogin}
							className="flex w-full justify-center rounded-md bg-blue-500 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-rose-600/60 focus:outline-none focus:ring-2 focus:ring-rose-500 focus:ring-offset-2"
						>
							Logg inn med Facebook
						</button>
					</div>
				</div>
			</div>
		</div>
	);
}
