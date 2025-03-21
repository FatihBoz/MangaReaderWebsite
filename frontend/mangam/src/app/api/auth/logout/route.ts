import { NextResponse } from 'next/server';
import { cookies } from 'next/headers';

export async function POST() {
  try {
    // Clear all auth related cookies
    const cookieStore = cookies();
    (await cookieStore).delete('user');
    (await cookieStore).delete('auth_token');
    (await cookieStore).delete('session');

    // Backend'e logout isteği gönder
    await fetch(`${process.env.NEXT_PUBLIC_API_URL}/users/logout`, {
      method: 'POST',
      credentials: 'include',
    });

    return NextResponse.json(
      { message: 'Logged out successfully' },
      { status: 200 }
    );
  } catch (error) {
    console.error('Logout error:', error);
    return NextResponse.json(
      { message: 'Logout failed' },
      { status: 500 }
    );
  }
}
