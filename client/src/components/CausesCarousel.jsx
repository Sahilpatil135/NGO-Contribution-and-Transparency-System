import React from 'react';
import img1 from "/img1.png";

// Example data for causes
const causes = [
//   { id: 1, name: "Urgent", img: "/images/urgent.jpg" }, add links for respective pages
  { id: 1, name: "Urgent", img: "/carousel_urgent.png" },
  { id: 2, name: "Strays", img: "carousel_strays.png" },
  { id: 3, name: "Elderly", img: "/carousel_elderly.png" },
  { id: 4, name: "Children", img: img1 },
  { id: 5, name: "Environmental", img: "/carousel_environmental.png" },
  { id: 6, name: "Specially-Abled", img: img1 },
  { id: 7, name: "Education", img: img1 },
  { id: 8, name: "Hunger", img: img1 },
  { id: 9, name: "Faith", img: img1 },
  { id: 10, name: "Health", img: img1 },
  { id: 11, name: "Poverty", img: img1 },
  { id: 12, name: "Women", img: img1 },
  { id: 13, name: "Arts & Culture", img: img1 },
  { id: 14, name: "Sports", img: img1 }
];

const CausesCarousel = () => {
  return (
    <section className="my-16 bg-gray-200">
      <div className="w-11/12 mx-auto py-8">
        <h1 className="text-4xl mb-4">Explore Causes</h1>
        <div className="overflow-x-auto mx-12">
          <div className="flex space-x-6 py-4">
            {causes.map((cause) => (
              <div
                key={cause.id}
                className="flex-shrink-0 w-36 md:w-48 bg-transparent cursor-pointer grayscale hover:grayscale-0 transition duration-300"
              >
                <img
                  src={cause.img}
                  alt={cause.name}
                  className="w-full h-32 md:h-30 object-cover p-1 rounded-md border-2 border-transparent hover:border-[#ff6200] transition-all duration-300"
                />
                <h3 className="p-2 text-center text-sm font-semibold hover: text-[#ff6200]">{cause.name}</h3>
              </div>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
};

export default CausesCarousel;
