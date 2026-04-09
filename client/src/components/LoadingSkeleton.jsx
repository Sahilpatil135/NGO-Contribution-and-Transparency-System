// Loading skeleton components for better UX during async operations

export const CardSkeleton = () => (
  <div className="border rounded-lg p-4 bg-white shadow-sm animate-pulse">
    <div className="h-48 bg-gray-200 rounded-lg mb-4"></div>
    <div className="h-6 bg-gray-200 rounded w-3/4 mb-2"></div>
    <div className="h-4 bg-gray-200 rounded w-full mb-2"></div>
    <div className="h-4 bg-gray-200 rounded w-5/6 mb-4"></div>
    <div className="flex justify-between items-center">
      <div className="h-4 bg-gray-200 rounded w-1/4"></div>
      <div className="h-8 bg-gray-200 rounded w-1/3"></div>
    </div>
  </div>
);

export const CampaignCardSkeleton = () => (
  <div className="border rounded-lg overflow-hidden bg-white shadow-sm animate-pulse">
    <div className="h-56 bg-gray-200"></div>
    <div className="p-4">
      <div className="h-6 bg-gray-200 rounded w-3/4 mb-3"></div>
      <div className="h-4 bg-gray-200 rounded w-full mb-2"></div>
      <div className="h-4 bg-gray-200 rounded w-5/6 mb-4"></div>
      <div className="h-2 bg-gray-200 rounded w-full mb-2"></div>
      <div className="flex justify-between items-center mb-3">
        <div className="h-4 bg-gray-200 rounded w-1/3"></div>
        <div className="h-4 bg-gray-200 rounded w-1/4"></div>
      </div>
      <div className="h-10 bg-gray-200 rounded"></div>
    </div>
  </div>
);

export const GridSkeleton = ({ count = 6, Component = CardSkeleton }) => (
  <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
    {Array.from({ length: count }).map((_, i) => (
      <Component key={i} />
    ))}
  </div>
);

export const ListSkeleton = ({ count = 5 }) => (
  <div className="space-y-4">
    {Array.from({ length: count }).map((_, i) => (
      <div key={i} className="border rounded-lg p-4 bg-white shadow-sm animate-pulse">
        <div className="flex items-start gap-4">
          <div className="w-20 h-20 bg-gray-200 rounded"></div>
          <div className="flex-1">
            <div className="h-5 bg-gray-200 rounded w-3/4 mb-2"></div>
            <div className="h-4 bg-gray-200 rounded w-full mb-2"></div>
            <div className="h-4 bg-gray-200 rounded w-5/6"></div>
          </div>
        </div>
      </div>
    ))}
  </div>
);

export const TableSkeleton = ({ rows = 5, cols = 4 }) => (
  <div className="overflow-x-auto">
    <table className="min-w-full divide-y divide-gray-200">
      <thead className="bg-gray-50">
        <tr>
          {Array.from({ length: cols }).map((_, i) => (
            <th key={i} className="px-6 py-3">
              <div className="h-4 bg-gray-200 rounded animate-pulse"></div>
            </th>
          ))}
        </tr>
      </thead>
      <tbody className="bg-white divide-y divide-gray-200">
        {Array.from({ length: rows }).map((_, rowIndex) => (
          <tr key={rowIndex}>
            {Array.from({ length: cols }).map((_, colIndex) => (
              <td key={colIndex} className="px-6 py-4">
                <div className="h-4 bg-gray-200 rounded animate-pulse"></div>
              </td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  </div>
);

export const ImageSkeleton = ({ className = "h-48 w-full" }) => (
  <div className={`bg-gray-200 rounded animate-pulse ${className}`}></div>
);

export const TextSkeleton = ({ lines = 3, className = "" }) => (
  <div className={`space-y-2 ${className}`}>
    {Array.from({ length: lines }).map((_, i) => (
      <div
        key={i}
        className={`h-4 bg-gray-200 rounded animate-pulse ${
          i === lines - 1 ? "w-3/4" : "w-full"
        }`}
      ></div>
    ))}
  </div>
);

export const ButtonSkeleton = ({ className = "h-10 w-32" }) => (
  <div className={`bg-gray-200 rounded animate-pulse ${className}`}></div>
);

export const AvatarSkeleton = ({ size = "h-12 w-12" }) => (
  <div className={`${size} bg-gray-200 rounded-full animate-pulse`}></div>
);

export const UpdateCardSkeleton = () => (
  <div className="border rounded-lg p-4 bg-white shadow-sm animate-pulse">
    <div className="flex items-start gap-3 mb-3">
      <AvatarSkeleton />
      <div className="flex-1">
        <div className="h-5 bg-gray-200 rounded w-1/2 mb-2"></div>
        <div className="h-3 bg-gray-200 rounded w-1/4"></div>
      </div>
    </div>
    <TextSkeleton lines={2} />
    <div className="grid grid-cols-3 gap-2 mt-4">
      <ImageSkeleton className="h-24" />
      <ImageSkeleton className="h-24" />
      <ImageSkeleton className="h-24" />
    </div>
  </div>
);

export const ProfileSkeleton = () => (
  <div className="max-w-4xl mx-auto p-6">
    <div className="flex items-center gap-6 mb-8 animate-pulse">
      <AvatarSkeleton size="h-24 w-24" />
      <div className="flex-1">
        <div className="h-8 bg-gray-200 rounded w-1/3 mb-3"></div>
        <div className="h-4 bg-gray-200 rounded w-1/2"></div>
      </div>
    </div>
    <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
      <TextSkeleton lines={4} />
      <TextSkeleton lines={4} />
    </div>
  </div>
);
