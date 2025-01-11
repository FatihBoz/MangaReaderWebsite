import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  images: {
      remotePatterns: [
          {
              protocol: 'https',
              hostname: 'placehold.co',
              port: '',
              pathname: '**',
          },
            {
                protocol: 'https',
                hostname: 'upload.wikimedia.org',
                port: '',
                pathname: '**',
            },{
                protocol: 'https',
                hostname: 'lh3.googleusercontent.com',
                port: '',
                pathname: '**',
            },
            
      ],
  },
};

export default nextConfig;
