import React from 'react';

interface HeroProps {
  title: string;
  subtitle?: string;
  gradient: string;
  height?: number;
}

const HeroSection: React.FC<HeroProps> = ({ title, subtitle, gradient, height }) => {
  return (
    <div
      className="relative"
      style={{
        height: height ?? 500,
      }}
    >
      <div
        className="absolute inset-0"
        style={{
          backgroundImage: gradient,
          opacity: 0.2,
        }}
      ></div>
      <div className="container mx-auto h-full flex items-center justify-center">
        <div className="text-center">
          <h1 className="text-4xl sm:text-5xl md:text-6xl lg:text-7xl font-bold text-black leading-tight">
            {title}
          </h1>
          {subtitle && (
            <h2 className="text-md sm:text-lg md:text-xl lg:text-2xl text-gray-800 mt-3">
              {subtitle}
            </h2>
          )}
        </div>
      </div>
    </div>
  );
};

export default HeroSection;
