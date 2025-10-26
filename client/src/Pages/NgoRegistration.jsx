 import React, { useState } from "react";

const NgoRegistration = () => {
  const [formData, setFormData] = useState({
    organization_name: "",
    registration_number: "",
    organization_type: "",
    about: "",
    website_url: "",
    address: "",
  });

  const [primaryContact, setPrimaryContact] = useState({
    name: "",
    role: "",
    email: "",
    phone: "",
  });

  const [documents, setDocuments] = useState({
    registration_certificate: null,
    pan_card: null,
    other_docs: null,
  });

  const [loading, setLoading] = useState(false);

  // handle input changes
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleContactChange = (e) => {
    const { name, value } = e.target;
    setPrimaryContact({ ...primaryContact, [name]: value });
  };

  const handleFileChange = (e) => {
    const { name, files } = e.target;
    setDocuments({ ...documents, [name]: files[0] });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // basic validation
    if (!formData.organization_name || !primaryContact.email) {
      alert("Please fill all required fields");
      return;
    }

    setLoading(true);

    const data = new FormData();
    data.append("organization_name", formData.organization_name);
    data.append("registration_number", formData.registration_number);
    data.append("organization_type", formData.organization_type);
    data.append("about", formData.about);
    data.append("website_url", formData.website_url);
    data.append("address", formData.address);

    // Contact details
    data.append("contact_name", primaryContact.name);
    data.append("contact_role", primaryContact.role);
    data.append("contact_email", primaryContact.email);
    data.append("contact_phone", primaryContact.phone);

    // Documents
    if (documents.registration_certificate)
      data.append("registration_certificate", documents.registration_certificate);
    if (documents.pan_card)
      data.append("pan_card", documents.pan_card);
    if (documents.other_docs)
      data.append("other_docs", documents.other_docs);

    try {
      const res = await fetch("http://localhost:8080/api/ngo/register", {
        method: "POST",
        body: data,
      });

      if (res.ok) {
        alert("✅ NGO registered successfully!");
        setFormData({
          organization_name: "",
          registration_number: "",
          organization_type: "",
          about: "",
          website_url: "",
          address: "",
        });
        setPrimaryContact({ name: "", role: "", email: "", phone: "" });
        setDocuments({ registration_certificate: null, pan_card: null, other_docs: null });
      } else {
        alert("❌ Error registering NGO");
      }
    } catch (error) {
      console.error(error);
      alert("Error connecting to server");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="bg-gray-100 py-12">
      <div className="max-w-5xl mx-auto bg-white rounded-lg shadow-md p-10">
        <h1 className="text-3xl font-bold text-[#3a0b2e] mb-8 text-center">
          Register Your NGO / Organization
        </h1>

        <form onSubmit={handleSubmit} className="space-y-8">

          {/* --- Organization Details --- */}
          <section>
            <h2 className="text-xl font-semibold text-[#ff6200] mb-4">
              Organization Details
            </h2>
            <div className="grid md:grid-cols-2 gap-6">
              <div>
                <label className="block text-sm font-medium mb-2">
                  Organization Name <span className="text-red-500">*</span>
                </label>
                <input
                  type="text"
                  name="organization_name"
                  value={formData.organization_name}
                  onChange={handleChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                  placeholder="Enter organization name"
                />
              </div>

              <div>
                <label className="block text-sm font-medium mb-2">
                  Registration Number
                </label>
                <input
                  type="text"
                  name="registration_number"
                  value={formData.registration_number}
                  onChange={handleChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                  placeholder="Enter registration number"
                />
              </div>

              <div>
                <label className="block text-sm font-medium mb-2">
                  Organization Type
                </label>
                <select
                  name="organization_type"
                  value={formData.organization_type}
                  onChange={handleChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                >
                  <option value="">Select type</option>
                  <option value="NGO">NGO</option>
                  <option value="Trust">Trust</option>
                  <option value="Society">Society</option>
                  <option value="Foundation">Foundation</option>
                </select>
              </div>

              <div>
                <label className="block text-sm font-medium mb-2">Website</label>
                <input
                  type="url"
                  name="website_url"
                  value={formData.website_url}
                  onChange={handleChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                  placeholder="https://example.org"
                />
              </div>
            </div>

            <div className="mt-4">
              <label className="block text-sm font-medium mb-2">About Organization</label>
              <textarea
                name="about"
                value={formData.about}
                onChange={handleChange}
                rows="3"
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                placeholder="Brief description about your organization"
              ></textarea>
            </div>

            <div className="mt-4">
              <label className="block text-sm font-medium mb-2">Address</label>
              <textarea
                name="address"
                value={formData.address}
                onChange={handleChange}
                rows="2"
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                placeholder="Enter address"
              ></textarea>
            </div>
          </section>

          {/* --- Primary Contact --- */}
          <section>
            <h2 className="text-xl font-semibold text-[#ff6200] mb-4">
              Primary Contact Details
            </h2>
            <div className="grid md:grid-cols-2 gap-6">
              <div>
                <label className="block text-sm font-medium mb-2">Name</label>
                <input
                  type="text"
                  name="name"
                  value={primaryContact.name}
                  onChange={handleContactChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                  placeholder="Full name of contact person"
                />
              </div>

              <div>
                <label className="block text-sm font-medium mb-2">Role</label>
                <input
                  type="text"
                  name="role"
                  value={primaryContact.role}
                  onChange={handleContactChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                  placeholder="e.g. Founder, Secretary, Manager"
                />
              </div>

              <div>
                <label className="block text-sm font-medium mb-2">Email</label>
                <input
                  type="email"
                  name="email"
                  value={primaryContact.email}
                  onChange={handleContactChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                  placeholder="example@email.com"
                />
              </div>

              <div>
                <label className="block text-sm font-medium mb-2">Phone</label>
                <input
                  type="tel"
                  name="phone"
                  value={primaryContact.phone}
                  onChange={handleContactChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                  placeholder="10-digit phone number"
                  maxLength="10"
                />
              </div>
            </div>
          </section>

          {/* --- Documents Upload --- */}
          <section>
            <h2 className="text-xl font-semibold text-[#ff6200] mb-4">
              Upload Documents
            </h2>
            <div className="grid md:grid-cols-3 gap-6">
              <div>
                <label className="block text-sm font-medium mb-2">
                  Registration Certificate
                </label>
                <input
                  type="file"
                  name="registration_certificate"
                  onChange={handleFileChange}
                  className="block w-full text-sm text-gray-700 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-white file:bg-[#ff6200] hover:file:bg-[#e45a00] cursor-pointer file:cursor-pointer"
                />
              </div>

              <div>
                <label className="block text-sm font-medium mb-2">PAN Card</label>
                <input
                  type="file"
                  name="pan_card"
                  onChange={handleFileChange}
                  className="block w-full text-sm text-gray-700 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-white file:bg-[#ff6200] hover:file:bg-[#e45a00] file:cursor-pointer cursor-pointer"
                />
              </div>

              <div>
                <label className="block text-sm font-medium mb-2">Other Documents</label>
                <input
                  type="file"
                  name="other_docs"
                  onChange={handleFileChange}
                  className="block w-full text-sm text-gray-700 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-white file:bg-[#ff6200] hover:file:bg-[#e45a00] file:cursor-pointer cursor-pointer"
                />
              </div>
            </div>
          </section>

          {/* --- Submit --- */}
          <div className="text-center">
            <button
              type="submit"
              disabled={loading}
              className="bg-[#ff6200] text-white font-semibold px-8 py-3 rounded-lg hover:bg-[#e45a00] transition disabled:opacity-50 cursor-pointer"
            >
              {loading ? "Submitting..." : "Register NGO"}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default NgoRegistration;
