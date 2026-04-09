// Enhanced form validation with better UX
import { useState } from 'react';

// Validation rules
export const ValidationRules = {
  required: (value) => (value && value.toString().trim() ? null : 'This field is required'),
  email: (value) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(value) ? null : 'Please enter a valid email address';
  },
  minLength: (min) => (value) => {
    return value && value.length >= min ? null : `Must be at least ${min} characters`;
  },
  maxLength: (max) => (value) => {
    return value && value.length <= max ? null : `Must be at most ${max} characters`;
  },
  min: (min) => (value) => {
    return value >= min ? null : `Must be at least ${min}`;
  },
  max: (max) => (value) => {
    return value <= max ? null : `Must be at most ${max}`;
  },
  pattern: (regex, message) => (value) => {
    return regex.test(value) ? null : message;
  },
  url: (value) => {
    try {
      new URL(value);
      return null;
    } catch {
      return 'Please enter a valid URL';
    }
  },
  phone: (value) => {
    const phoneRegex = /^[0-9]{10,15}$/;
    return phoneRegex.test(value.replace(/\D/g, '')) ? null : 'Please enter a valid phone number';
  },
};

// Form field component with validation
export const FormField = ({
  label,
  name,
  type = 'text',
  value,
  onChange,
  error,
  touched,
  onBlur,
  placeholder,
  required = false,
  helpText,
  className = '',
  ...props
}) => {
  return (
    <div className={`mb-4 ${className}`}>
      <label htmlFor={name} className="block text-sm font-medium text-gray-700 mb-1">
        {label}
        {required && <span className="text-red-500 ml-1">*</span>}
      </label>
      <input
        type={type}
        id={name}
        name={name}
        value={value}
        onChange={onChange}
        onBlur={onBlur}
        placeholder={placeholder}
        className={`w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 transition ${
          error && touched
            ? 'border-red-500 focus:ring-red-500 focus:border-red-500'
            : 'border-gray-300 focus:ring-blue-500 focus:border-blue-500'
        }`}
        aria-invalid={error && touched ? 'true' : 'false'}
        aria-describedby={error && touched ? `${name}-error` : undefined}
        {...props}
      />
      {helpText && !error && <p className="mt-1 text-sm text-gray-500">{helpText}</p>}
      {error && touched && (
        <p id={`${name}-error`} className="mt-1 text-sm text-red-600 flex items-center gap-1">
          <svg className="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
            <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
          </svg>
          {error}
        </p>
      )}
    </div>
  );
};

// Textarea field component
export const TextAreaField = ({
  label,
  name,
  value,
  onChange,
  error,
  touched,
  onBlur,
  placeholder,
  required = false,
  rows = 4,
  helpText,
  maxLength,
  className = '',
  ...props
}) => {
  return (
    <div className={`mb-4 ${className}`}>
      <label htmlFor={name} className="block text-sm font-medium text-gray-700 mb-1">
        {label}
        {required && <span className="text-red-500 ml-1">*</span>}
      </label>
      <textarea
        id={name}
        name={name}
        value={value}
        onChange={onChange}
        onBlur={onBlur}
        placeholder={placeholder}
        rows={rows}
        maxLength={maxLength}
        className={`w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 transition ${
          error && touched
            ? 'border-red-500 focus:ring-red-500 focus:border-red-500'
            : 'border-gray-300 focus:ring-blue-500 focus:border-blue-500'
        }`}
        aria-invalid={error && touched ? 'true' : 'false'}
        aria-describedby={error && touched ? `${name}-error` : undefined}
        {...props}
      />
      <div className="flex justify-between items-start mt-1">
        <div className="flex-1">
          {helpText && !error && <p className="text-sm text-gray-500">{helpText}</p>}
          {error && touched && (
            <p id={`${name}-error`} className="text-sm text-red-600 flex items-center gap-1">
              <svg className="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
              </svg>
              {error}
            </p>
          )}
        </div>
        {maxLength && (
          <p className="text-sm text-gray-400">
            {value?.length || 0}/{maxLength}
          </p>
        )}
      </div>
    </div>
  );
};

// Custom hook for form validation
export const useFormValidation = (initialValues, validationRules) => {
  const [values, setValues] = useState(initialValues);
  const [errors, setErrors] = useState({});
  const [touched, setTouched] = useState({});
  const [isSubmitting, setIsSubmitting] = useState(false);

  const validateField = (name, value) => {
    const rules = validationRules[name];
    if (!rules) return null;

    for (const rule of rules) {
      const error = rule(value);
      if (error) return error;
    }
    return null;
  };

  const validateAll = () => {
    const newErrors = {};
    Object.keys(validationRules).forEach((name) => {
      const error = validateField(name, values[name]);
      if (error) newErrors[name] = error;
    });
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    const newValue = type === 'checkbox' ? checked : value;
    
    setValues((prev) => ({ ...prev, [name]: newValue }));
    
    // Validate on change if field has been touched
    if (touched[name]) {
      const error = validateField(name, newValue);
      setErrors((prev) => ({ ...prev, [name]: error }));
    }
  };

  const handleBlur = (e) => {
    const { name } = e.target;
    setTouched((prev) => ({ ...prev, [name]: true }));
    
    // Validate on blur
    const error = validateField(name, values[name]);
    setErrors((prev) => ({ ...prev, [name]: error }));
  };

  const resetForm = () => {
    setValues(initialValues);
    setErrors({});
    setTouched({});
    setIsSubmitting(false);
  };

  return {
    values,
    errors,
    touched,
    isSubmitting,
    setIsSubmitting,
    handleChange,
    handleBlur,
    validateAll,
    resetForm,
    setValues,
  };
};

// Success message component
export const SuccessMessage = ({ message }) => (
  <div className="mb-4 p-4 bg-green-50 border border-green-200 rounded-md flex items-center gap-2">
    <svg className="w-5 h-5 text-green-600" fill="currentColor" viewBox="0 0 20 20">
      <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clipRule="evenodd" />
    </svg>
    <p className="text-sm text-green-800">{message}</p>
  </div>
);

// Error message component
export const ErrorMessage = ({ message }) => (
  <div className="mb-4 p-4 bg-red-50 border border-red-200 rounded-md flex items-center gap-2">
    <svg className="w-5 h-5 text-red-600" fill="currentColor" viewBox="0 0 20 20">
      <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
    </svg>
    <p className="text-sm text-red-800">{message}</p>
  </div>
);
