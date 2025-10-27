import { useState, useEffect } from "react"
import { useParams } from "react-router-dom";
import { apiRequest, API_ENDPOINTS } from "../config/api";
import Card from "../components/Card";

const DonationTypePage = () => {
  const { category, slug } = useParams();

  const [causeCategory, setCauseCategory] = useState({})
  const [loadingCategory, setLoadingCategory] = useState(true)
  const [causes, setCauses] = useState([])
  const [loadingCauses, setLoadingCauses] = useState(true)

  useEffect(() => {
    const fetchCategoryData = async () => {
      let apiEndpoint

      if (category == "aidType") {
        apiEndpoint = API_ENDPOINTS.GET_ALL_AID_TYPES
      } else {
        apiEndpoint = API_ENDPOINTS.GET_ALL_DOMAINS
      }

      const categoryResult = await apiRequest(apiEndpoint)

      if (categoryResult.success && categoryResult.data) {
        categoryResult.data.forEach(category => {
          if (slugFromName(category.name) == slug) {
            setCauseCategory(category)
          }
        });

        setLoadingCategory(false)
      }
    }

    fetchCategoryData()
  }, [slug, category])

  useEffect(() => {
    const fetchCauseData = async () => {
      let apiEndpoint

      if (category == "domain") {
        apiEndpoint = API_ENDPOINTS.GET_CAUSES_BY_DOMAIN
      } else if (category == "aidType") {
        apiEndpoint = API_ENDPOINTS.GET_CAUSES_BY_AID_TYPE
      } else {
        apiEndpoint = API_ENDPOINTS.GET_CAUSES
      }

      const causesResult = await apiRequest(`${apiEndpoint}/${causeCategory.id}`)

      if (causesResult.success && causesResult.data) {
        setCauses(causesResult.data)
        setLoadingCauses(false)
      }
    }

    if (causeCategory.id) {
      fetchCauseData()
    }
  }, [causeCategory])

  const slugFromName = (name => {
    return name.toLowerCase().replaceAll(" ", "-");
  })

  return (
    <>
      {!loadingCategory ? (
        <div className="max-w-7xl mx-auto py-12 px-4">
          <img src={causeCategory.icon_url} alt={causeCategory.name} className="w-full h-64 object-cover rounded-lg mb-12" />
          <h1 className="text-center text-4xl font-bold text-[#3a0b2e] mb-6">{causeCategory.name}</h1>
          <p className="text-center text-gray-700 mb-12">
            Explore campaigns, projects, and opportunities related to {causeCategory.name.toLowerCase()}.
          </p>

          {/* Replace below with dynamic cards or projects */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {!loadingCauses ? (
              causes.map((cause, idx) => (
                <Card key={idx} cause={cause} />
              ))
            ) : (
              <h1>LOADING</h1>
            )}
          </div>
        </div >
      ) : (<h1>LOADING</h1>)
      };
    </>
  )
};

export default DonationTypePage;
