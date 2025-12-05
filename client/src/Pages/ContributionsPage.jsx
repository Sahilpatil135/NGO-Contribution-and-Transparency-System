import { useState, useEffect } from "react";
import img1 from "/makeContributionbanner.jpg"
import img2 from "/banner_main.jpeg"
import CausesCarousel from '../components/CausesCarousel';
import Card from '../components/Card';
import { Link } from 'react-router-dom';
import { apiRequest, API_ENDPOINTS } from "../config/api";

const ContributionsPage = () => {

    const [aidTypes, setAidTypes] = useState([])
    const [loadingAidTypes, setLoadingAidTypes] = useState(true)
    const [loadingCauses, setLoadingCauses] = useState(true)
    const [causes, setCauses] = useState([])

    useEffect(() => {
        const fetchAidTypesData = async () => {
            const aidsResult = await apiRequest(API_ENDPOINTS.GET_ALL_AID_TYPES)
            if (aidsResult.success && aidsResult.data) {
                setAidTypes(aidsResult.data.slice(0, 5))
                setLoadingAidTypes(false)
            }
        }

        const fetchCausesData = async () => {
            const causeResult = await apiRequest(`${API_ENDPOINTS.GET_CAUSES}/?limit=3`)

            if (causeResult.success && causeResult.data) {
                setCauses(causeResult.data)
                setLoadingCauses(false)
            }
        }

        fetchAidTypesData()
        fetchCausesData()
    }, [])


    useEffect(() => {
    }, [])

    const slugFromName = (name => {
        return name.toLowerCase().replaceAll(" ", "-");
    })

    return (
        <div className="w-full">
            {/* Hero section  */}
            <section>
                <img src={img2} alt="..." className="w-full" />
            </section>

            {/* Types of Donations section  */}
            <section className="mt-30">
                <div className="w-11/12 flex flex-col items-center justify-between mx-auto my-6">
                    <h1 className="text-5xl font-bold mb-16">Where Do <span className="text-[#ff6200] italic">YOU</span> Want To <span className="text-[#ff6200] italic">Contribute</span> ?</h1>
                    <div className="grid grid-cols-1 md:grid-cols-5 gap-8 w-full my-8">
                        {/* will replace by clickable div's with links to respective pages */}
                        {/* <div className="flex flex-col"><img src={img2} alt=".." /> Monetary Donations</div>
                        <div className="flex flex-col"><img src={img2} alt=".." /> Voluntering</div>
                        <div className="flex flex-col"><img src={img2} alt=".." /> Blood & Organ Donations</div>
                        <div className="flex flex-col"><img src={img2} alt=".." /> Goods & Resources Donation</div>
                        <div className="flex flex-col"><img src={img2} alt=".." /> Environmental Support</div> */}

                        {!loadingAidTypes ? aidTypes.map((aidType, idx) => (
                            <Link key={idx} to={`/makecontribution/aidType/${slugFromName(aidType.name)}`} className="flex flex-col items-center cursor-pointer">
                                <img src={aidType.icon_url} alt={aidType.name} className="rounded-lg" />
                                <p className="mt-2 font-semibold text-[#3a0b2e]">{aidType.name}</p>
                            </Link>)) : (
                            <h1>loading</h1>
                        )}


                    </div>
                </div >
            </section >

            {/* Explore NGO's  */}
            {/* <section className="my-16 bg-gray-200">
            <div className="w-11/12 flex flex-col justify-between mx-auto py-16">
                <h1 className="text-4xl pl-16">Explore Causes</h1>
                
            </div>
        </section> */}

            {/* Types of Causes */}
            <section>
                <CausesCarousel />
            </section>

            {/* Dynamically causes card will get displayed on selecting type of cause  */}
            <section className="w-11/12 mx-auto my-20 items-center text-center space-y-6">
                <div className="grid grid-cols-1 md:grid-cols-3 gap-12 w-full my-8">
                    {!loadingCauses ? (
                        causes.map((cause, idx) => (
                            <Card key={idx} cause={cause} />
                        ))
                    ) : (
                        <h1>LOADING</h1>
                    )}
                </div>
            </section>
        </div >
    )
}

export default ContributionsPage
