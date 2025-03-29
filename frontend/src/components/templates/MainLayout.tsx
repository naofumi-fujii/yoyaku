'use client';

import React from 'react';
import Typography from '../atoms/Typography';

interface MainLayoutProps {
  title?: string;
  children: React.ReactNode;
}

export default function MainLayout({ title, children }: MainLayoutProps) {
  return (
    <div className="min-h-screen bg-gray-50">
      {/* ヘッダー */}
      <header className="bg-white shadow-sm">
        <div className="container mx-auto px-4 py-4">
          <Typography variant="h1">
            {title || '予約システム'}
          </Typography>
        </div>
      </header>
      
      {/* メインコンテンツ */}
      <main className="container mx-auto px-4 py-6">
        {children}
      </main>
      
      {/* フッター */}
      <footer className="bg-white border-t border-gray-200 mt-10">
        <div className="container mx-auto px-4 py-4 text-center">
          <Typography variant="caption" color="text-gray-500">
            &copy; 2024 予約システム
          </Typography>
        </div>
      </footer>
    </div>
  );
}