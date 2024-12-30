import Link from 'next/link';

export default function Unauthorized() {
  return (
    <div>
      <h2>Unauthorized</h2>
      <p>Please login to view this content.</p>
      <Link href="/">Return Home</Link>
    </div>
  );
}
