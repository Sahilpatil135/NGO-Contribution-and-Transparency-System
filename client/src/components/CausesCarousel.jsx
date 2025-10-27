import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { apiRequest, API_ENDPOINTS } from '../config/api';

const CausesCarousel = () => {

  const [domains, setDomains] = useState([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      const aidsResult = await apiRequest(API_ENDPOINTS.GET_ALL_DOMAINS)
      if (aidsResult.success && aidsResult.data) {
        setDomains(aidsResult.data)
        setLoading(false)
      }
    }

    fetchData()
  }, [])

  const slugFromName = (name => {
    return name.toLowerCase().replaceAll(" ", "-");
  })

  return (
    <section className="my-16 bg-gray-200">
      <div className="w-11/12 mx-auto py-8">
        <h1 className="text-4xl mb-4">Explore Causes</h1>
        <div className="overflow-x-auto mx-12">
          <div className="flex space-x-6 py-4">
            {!loading ? (
              domains.map((domain, idx) => (
                <Link key={idx} to={`/makecontribution/domain/${slugFromName(domain.name)}`} className="flex flex-col items-center cursor-pointer">
                  <div
                    key={idx}
                    className="flex-shrink-0 w-36 md:w-48 bg-transparent cursor-pointer grayscale hover:grayscale-0 transition duration-300"
                  >
                    <img
                      src={domain.icon_url}
                      alt={domain.name}
                      className="w-full h-32 md:h-30 object-cover p-1 rounded-md border-2 border-transparent hover:border-[#ff6200] transition-all duration-300"
                    />
                    <h3 className="p-2 text-center text-sm font-semibold hover: text-[#ff6200]">{domain.name}</h3>
                  </div>
                </Link>
              ))) : (
              <h1>LOADING</h1>
            )}
          </div>
        </div>
      </div>
    </section>
  );
};

export default CausesCarousel;
