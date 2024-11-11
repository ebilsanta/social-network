import jwt from 'jsonwebtoken';
import NextAuth, { AuthOptions } from 'next-auth';
import type { JWT, JWTDecodeParams, JWTEncodeParams } from 'next-auth/jwt';
import GoogleProvider from 'next-auth/providers/google';

export const authOptions: AuthOptions = {
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID!,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET!,
    }),
  ],
  callbacks: {
    session: async ({ session, token }) => {
      if (session?.user) {
        session.user.id = token.sub;
        session.user.accessToken = token.accessToken as string;
      }
      return session;
    },
    async jwt({ token, account }) {
      if (account) {
        token.iss = process.env.NEXTAUTH_ISSUER;
        token.exp = account.expires_at;
      }
      return token;
    },
  },
  secret: process.env.NEXTAUTH_SECRET,
  jwt: {
    async encode(params: JWTEncodeParams): Promise<string> {
      return jwt.sign(params.token!, params.secret, {
        algorithm: 'HS256',
      });
    },
    async decode(params: JWTDecodeParams): Promise<JWT | null> {
      return jwt.verify(params.token!, params.secret) as JWT;
    },
  },
};

const handler = NextAuth(authOptions);

export { handler as GET, handler as POST };
