"use client";

import React, {ReactNode, useEffect, useMemo, useRef, useState} from "react";
import {AlertCircle, ArrowLeft, CheckCircle2, Eye, EyeOff, Fingerprint, Mail, RefreshCw,} from "lucide-react";
import {Button as StyledButton} from "@/components/ui/button";
import {Card, CardContent, CardDescription, CardHeader, CardTitle,} from "@/components/ui/card";
import {Input as StyledInput} from "@/components/ui/input";
import {Label} from "@/components/ui/label";
import {Tabs, TabsContent, TabsList, TabsTrigger} from "@/components/ui/tabs";
import {Separator} from "@/components/ui/separator";
import {Checkbox} from "@/components/ui/checkbox";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue,} from "@/components/ui/select";
import {AuthView, FormField, FrankConfig, Link} from "./context";
import {useFrank} from "./hooks";
import {useUrlSearchParams} from "../hooks/use-url-search-params";
import {cn} from "@/lib/utils";
import {
	authForgotPassword,
	authLogin,
	authRegister,
	authResetPassword,
	LoginResponse2,
	LoginResponse3,
	passkeysLoginBegin,
	passwordlessMagicLink,
} from "@frank-auth/sdk";
import {Alert, AlertDescription, AlertTitle} from "@/components/ui/alert";
import {InputOTP, InputOTPGroup, InputOTPSlot,} from "@/components/ui/input-otp";
import {REGEXP_ONLY_DIGITS_AND_CHARS} from "input-otp";
import {setTokenData} from "@/utils/token";

export interface FrankProps extends FrankConfig {
	useProviderConfig?: boolean;
}

export function FrankUIKit({
	logo,
	title,
	titleAlign,
	description,
	oauthProviders,
	onLogin,
	onSignup,
	onPasswordless,
	onPasskey,
	onMfa,
	onForgotPassword,
	onVerifyOtp,
	onResendVerification,
							   onResetPassword,
	supportedMethods,
	showTabs,
	availableTabs,
	initialView,
	signupFields,
	theme,
	links,
	verification,
	useProviderConfig = true,
	components,
	api,
	cssClasses,
}: FrankProps) {
	// Get configuration from context if useProviderConfig is true
	const { config: providerConfig, ...frank } = useFrank();
	const searchParams = useUrlSearchParams();

	// Merge the provider config with props, prioritizing props
	const config = useProviderConfig
		? {
				// Start with provider config as the base
				...providerConfig,
				// Override with any directly provided props that aren't undefined
				...(logo !== undefined && { logo }),
				...(title !== undefined && { title }),
				...(titleAlign !== undefined && { titleAlign }),
				...(description !== undefined && { description }),
				...(oauthProviders !== undefined && { oauthProviders }),
				...(onLogin !== undefined && { onLogin }),
				...(onSignup !== undefined && { onSignup }),
				...(onPasswordless !== undefined && { onPasswordless }),
				...(onResendVerification !== undefined && { onResendVerification }),
				...(onPasskey !== undefined && { onPasskey }),
				...(onMfa !== undefined && { onMfa }),
				...(onForgotPassword !== undefined && { onForgotPassword }),
				...(onVerifyOtp !== undefined && { onVerifyOtp }),
			...(onResetPassword !== undefined && { onResetPassword }),
				...(supportedMethods !== undefined && { supportedMethods }),
				...(showTabs !== undefined && { showTabs }),
				...(availableTabs !== undefined && { availableTabs }),
				...(initialView !== undefined && { initialView }),
				...(signupFields !== undefined && { signupFields }),
				...(theme !== undefined && { theme }),
				...(links !== undefined && { links }),
				...(verification !== undefined && { verification }),
				...(cssClasses !== undefined && { cssClasses }),
				...(components !== undefined && { components }),
				...(api !== undefined && { api }),
			}
		: {
				// If not using provider config, just use the props directly
				logo,
				title,
				titleAlign,
				description,
				oauthProviders,
				onLogin,
				onSignup,
				onPasswordless,
				onPasskey,
				onMfa,
				onForgotPassword,
				onVerifyOtp,
			onResetPassword,
				supportedMethods,
				onResendVerification,
				showTabs,
				availableTabs,
				initialView,
				signupFields,
				theme,
				links,
				verification,
				api,
				components,
				cssClasses,
			};

	// Set default values for required props (rest of the component remains the same)
	const {
		logo: configLogo,
		title: configTitle = "Welcome Back",
		titleAlign: configTitleAlign = "center",
		description: configDescription = "Sign in to your account",
		oauthProviders: configOauthProviders = [],
		onLogin: configOnLogin,
		onSignup: configOnSignup,
		onPasswordless: configOnPasswordless,
		onPasskey: configOnPasskey,
		onMfa: configOnMfa,
		onForgotPassword: configOnForgotPassword,
		onVerifyOtp: configOnVerifyOtp,
		onResendVerification: configOnResendVerification,
		onResetPassword: configOnResetPassword,
		supportedMethods: configSupportedMethods = [
			"password",
			"passwordless",
			"passkey",
		],
		api: configApi,
		showTabs: configShowTabs = true,
		availableTabs: configAvailableTabs = ["login", "signup"],
		initialView: configInitialView = "login",
		signupFields: configSignupFields = [
			{
				name: "name",
				label: "Name",
				type: "text",
				placeholder: "John Doe",
				required: true,
				row: 1,
				width: "half",
			},
			{
				name: "name",
				label: "Name",
				type: "text",
				placeholder: "John Doe",
				required: true,
				row: 1,
				width: "half",
			},
			{
				name: "email",
				label: "Email",
				type: "email",
				placeholder: "name@example.com",
				required: true,
			},
			{ name: "password", label: "Password", type: "password", required: true },
		],
		links: configLinks = {},
		verification: configVerification = {
			input: "otp",
			codeLength: 6,
		},
		theme: configThemeTemp = {
			primaryColor: "bg-primary",
			backgroundColor: "bg-background",
			textColor: "text-foreground",
			borderRadius: "rounded-md",
		},
		components: configComponents = {},
		cssClasses: configCssClasses = {},
	} = config;

	const Input = configComponents?.Input ?? StyledInput;
	const Button = configComponents?.Button ?? StyledButton;

	const [activeTab, setActiveTab] = useState<"login" | "signup">(
		configAvailableTabs.includes(configInitialView)
			? (configInitialView as "login" | "signup")
			: "login",
	);
	const [error, setError] = useState<string | null>(null);
	const [currentView, setCurrentView] = useState<AuthView>(configInitialView);
	const [loginStep, setLoginStep] = useState<number>(1);
	const [email, setEmail] = useState(searchParams.get("email") || "");
	const [password, setPassword] = useState("");
	const [showPassword, setShowPassword] = useState(false);
	const [mfaCode, setMfaCode] = useState("");
	const [otpCode, setOtpCode] = useState("");

	// Reset password states
	const [resetPasswordToken, setResetPasswordToken] = useState<string | null>(null);
	const [newPassword, setNewPassword] = useState("");
	const [confirmPassword, setConfirmPassword] = useState("");


	// Resend verification cooldown state
	const [cooldownRemaining, setCooldownRemaining] = useState(0);
	const cooldownTimerRef = useRef<NodeJS.Timeout | null>(null);

	// Status messages
	const [success, setSuccess] = useState("");
	const [isLoading, setIsLoading] = useState(false);

	// Dynamic form state for signup
	const [signupFormData, setSignupFormData] = useState<Record<string, any>>(
		() => {
			const initialData: Record<string, any> = {};
			configSignupFields.forEach((field) => {
				initialData[field.name] = field.type === "checkbox" ? false : "";
			});
			return initialData;
		},
	);
	const configTheme = useMemo(() => {
		return frank.getTheme(configThemeTemp);
	}, [configThemeTemp, frank]);

	// Apply theme styles
	const cardClassName = `w-full py-6 max-w-xl  min-w-md mx-auto ${configTheme.borderRadius || "rounded-md"} overflow-hidden ${configCssClasses?.card ?? ""}`;
	const buttonClassName = `${configTheme.primaryColor || "bg-primary"} ${configTheme.borderRadius || "rounded-md"}`;
	const inputClassName = `${configTheme.borderRadius || "rounded-md"}`;
	const cardStyle = configTheme.backgroundColor
		? { backgroundColor: configTheme.backgroundColor.replace("bg-", "") }
		: {};
	const textStyle = configTheme.textColor
		? { color: configTheme.textColor.replace("text-", "") }
		: {};
	const titleClasses = cn(
		"text-3xl font-bold",
		{
			"text-center": configTitleAlign === "center",
			"text-left": configTitleAlign === "left",
			"text-right": configTitleAlign === "right",
		},
		configCssClasses?.title,
	);
	const descriptionClasses = cn(
		"text-muted-foreground ",
		{
			"text-center": configTitleAlign === "center",
			"text-left": configTitleAlign === "left",
			"text-right": configTitleAlign === "right",
		},
		configCssClasses?.desc,
	);

	const codeInputLength = useMemo(() => {
		return new Array(configVerification.codeLength ?? 6).fill(0);
	}, [configVerification.codeLength]);

	// Clean up cooldown timer on unmount
	useEffect(() => {
		return () => {
			if (cooldownTimerRef.current) {
				clearInterval(cooldownTimerRef.current);
			}
		};
	}, []);

	// Check for token in URL
	useEffect(() => {
		// If we're in reset-password view, check for token
		if (currentView === "reset-password") {
			const token = searchParams?.get('token');

			if (token) {
				setResetPasswordToken(token);
			} else {
				// No token found, but we're in reset-password view
				setResetPasswordToken(null);
			}
		}
	}, [searchParams, currentView]);

	// Check if initialView is reset-password and set token if available
	useEffect(() => {
		if (configInitialView === "reset-password") {
			const token = searchParams?.get('token');
			if (token) {
				setResetPasswordToken(token);
			}
		}
	}, [configInitialView, searchParams]);

	const renderLogo = (): ReactNode => {
		if (!configLogo) return null;

		// If it's already JSX, render it directly
		return configLogo;
	};

	const resetState = () => {
		setLoginStep(1);
		setPassword("");
		setMfaCode("");
		setOtpCode("");
		setError(null);
		setSuccess("");
		setShowPassword(false);
		setNewPassword("");
		setConfirmPassword("");
	};

	const handleTabChange = (value: string) => {
		setActiveTab(value as "login" | "signup");
		resetState();
		setCurrentView(value as AuthView);
	};

	const handleEmailSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		setLoginStep(2);
	};

	const renderAlertMessages = () => (
		<div className="space-y-4 py-2">
			{error && (
				<Alert variant="destructive">
					<AlertCircle className="h-4 w-4" />
					<AlertTitle>Error</AlertTitle>
					<AlertDescription>{error}</AlertDescription>
				</Alert>
			)}

			{success && (
				<Alert
					variant="default"
					className="border-green-500 bg-green-500/10 text-green-500"
				>
					<CheckCircle2 className="h-4 w-4" />
					<AlertTitle>Success</AlertTitle>
					<AlertDescription>{success}</AlertDescription>
				</Alert>
			)}
		</div>
	);

	const postLoginOrSignup = async (rsp: LoginResponse3 | LoginResponse2) => {
		let href;

		if (rsp.requiresVerification) {
			if (configLinks?.verify) {
				href = `${configLinks.verify.url}?email=${email}&token=${rsp.token}`;
				href += `&redirect_url=${window.location.href}`;
				href += `&method=${rsp.verificationMethod}`;
				href += `&vid=${rsp.verificationId}`;
				window.location.href = href;
			} else {
				setCurrentView("verify-otp");
			}
			return;
		}

		if (rsp.message) {
			setSuccess(rsp.message);
		}

		if (rsp.mfa_required) {
			if (configLinks?.mfa) {
				href = `${configLinks.mfa.url}?email=${email}&token=${rsp.token}`;
				href += `&redirect_url=${window.location.href}`;
				href += `&method=${rsp.verificationMethod}`;
				href += `&vid=${rsp.verificationId}`;
				window.location.href = href;
			} else {
				setCurrentView("mfa");
			}
			return;
		}

		if (rsp.refresh_token && rsp.token && rsp.token.length > 0) {
			setTokenData({
				token: rsp.token,
				refreshToken: rsp.refresh_token,
				expiresAt: Number(rsp.expires_at),
			});
		}
	};

	const handleLogin = async (e: React.FormEvent) => {
		e.preventDefault();

		setIsLoading(true);
		setError(null);
		setError(null);

		try {
			const rsp = await authLogin({
				body: {
					email,
					password,
					organization_id: configApi?.projectId,
				},
			});

			if (rsp.error) {
				setError(rsp.error.message);
				return;
			}

			await configOnLogin?.({ email, password });
			await postLoginOrSignup(rsp.data);
			navigateToHome();
			resetState();
		} catch (error) {
			console.error("Login error:", error);
		} finally {
			setIsLoading(false);
		}
	};

	const navigateToLogin = () => {
		setTimeout(() => {
			if (configLinks?.login) {
				window.location.href = `${configLinks.login.url}?redirect_url=${window.location.href}`;
			} else {
				setCurrentView("login");
			}
		}, 500);
	};


	const navigateToHome = () => {
		setTimeout(() => {
			const red = configLinks?.redirectAfterLogin ?? searchParams?.get('redirect') ?? searchParams?.get('redirect_url') ?? searchParams?.get('redirect_after_login') ?? window.location.host;
			window.location.replace(red);
		}, 500);
	};

	const handleSignup = async (e: React.FormEvent) => {
		e.preventDefault();

		setIsLoading(true);
		setError(null);
		setError(null);

		try {
			const firstNameField = configSignupFields.find(
				(value) => value.isFirstName,
			);
			const lastNameField = configSignupFields.find(
				(value) => value.isLastName,
			);
			const emailField = configSignupFields.find(
				(value) => value.type === "email" || value.isEmail,
			);
			const passwordField = configSignupFields.find(
				(value) => value.type === "password",
			);

			const rsp = await authRegister({
				body: {
					email: signupFormData[emailField?.name || "email"],
					password: signupFormData[passwordField?.name || "password"],
					organization_id: configApi?.projectId,
					first_name: signupFormData[firstNameField?.name || "first_name"],
					last_name: signupFormData[lastNameField?.name || "last_name"],
					metadata: signupFormData,
				},
			});

			if (rsp.error) {
				setError(rsp.error.message);
				return;
			}

			await configOnSignup?.(signupFormData);

			await postLoginOrSignup(rsp.data);

			navigateToLogin();
			resetState();
		} catch (error) {
			console.error("Signup error:", error);
		} finally {
			setIsLoading(false);
		}
	};

	const handlePasswordless = async () => {
		if (!configOnPasswordless) return;

		setIsLoading(true);
		setError(null);
		setError(null);

		try {
			await configOnPasswordless(email);

			const rsp = await passwordlessMagicLink({
				// @ts-ignore
				body: {
					email,
				},
			});

			await configOnPasswordless?.(email);

			if (rsp.error) {
				setError(rsp.error.message);
				return;
			}
		} catch (error) {
			console.error("Passwordless error:", error);
		} finally {
			setIsLoading(false);
		}
	};

	const handlePasskey = async () => {
		setIsLoading(true);
		setError(null);
		setError(null);

		try {
			const rsp = await passkeysLoginBegin({
				// @ts-ignore
				body: {
					email,
				},
			});

			await configOnPasskey?.();

			if (rsp.error) {
				setError(rsp.error.message);
				return;
			}
		} catch (error) {
			console.error("Passkey error:", error);
		} finally {
			setIsLoading(false);
		}
	};

	const handleMfa = async (e: React.FormEvent) => {
		e.preventDefault();
		if (!configOnMfa) return;

		setIsLoading(true);
		setError(null);
		setError(null);

		try {
			await configOnMfa(mfaCode);
			setCurrentView("login");
			resetState();
		} catch (error) {
			console.error("MFA error:", error);
		} finally {
			setIsLoading(false);
		}
	};

	const handleForgotPassword = async (e: React.FormEvent) => {
		e.preventDefault();

		setIsLoading(true);
		setError(null);
		setError(null);

		try {
			const rsp = await authForgotPassword({
				body: {
					email,
				},
				query: {
					redirect_url:  configLinks?.resetPassword?.url ?? window.location.host,
				},
			});

			await configOnForgotPassword?.(email);

			if (rsp.error) {
				setError(rsp.error.message);
				return;
			}

			setSuccess(rsp.data.message);
		} catch (error) {
			console.error("Forgot password error:", error);
		} finally {
			setIsLoading(false);
		}
	};

	const handleResetPassword = async (e: React.FormEvent) => {
		e.preventDefault();

		setIsLoading(true);
		setError(null);

		// Validate passwords match
		if (newPassword !== confirmPassword) {
			setError("Passwords do not match");
			setIsLoading(false);
			return;
		}

		try {
			if (!resetPasswordToken) {
				setError("Reset token is missing or invalid");
				setIsLoading(false);
				return;
			}

			const rsp = await authResetPassword({
				body: {
					new_password: newPassword,
					token: resetPasswordToken,
				},
			});

			await configOnResetPassword?.(newPassword, resetPasswordToken);

			if (rsp.error) {
				setError(rsp.error.message);
				return;
			}

			setSuccess("Your password has been reset successfully");

			// Redirect to login after short delay
			setTimeout(() => {
				navigateToLogin();
			}, 3000);

		} catch (error) {
			console.error("Reset password error:", error);
			setError("An error occurred while resetting your password");
		} finally {
			setIsLoading(false);
		}
	};


	const handleVerifyOtp = async (e: React.FormEvent) => {
		e.preventDefault();
		setIsLoading(true);
		setError(null);
		setError(null);

		try {
			const rsp = await frank.verifyEmail({
				email,
				otp: otpCode,
				method: "otp",
			});

			await configOnVerifyOtp?.(email, otpCode);

			if (rsp.error) {
				setError(rsp.error.message);
				return;
			}

			navigateToLogin();
			resetState();
		} catch (error) {
			console.error("OTP verification error:", error);
		} finally {
			setIsLoading(false);
		}
	};

	const handleResendVerification = async () => {
		if (cooldownRemaining > 0) return;

		setIsLoading(true);
		try {
			const rsp = await frank.resendVerification({
				email,
				verification_type: "otp",
			});

			await configOnResendVerification?.(email);

			// Start cooldown timer
			setCooldownRemaining(frank.resendCooldown);

			if (cooldownTimerRef.current) {
				clearInterval(cooldownTimerRef.current);
			}

			cooldownTimerRef.current = setInterval(() => {
				setCooldownRemaining((prev) => {
					if (prev <= 1) {
						if (cooldownTimerRef.current) {
							clearInterval(cooldownTimerRef.current);
							cooldownTimerRef.current = null;
						}
						return 0;
					}
					return prev - 1;
				});
			}, 1000);
		} catch (error) {
			console.error("Resend verification error:", error);
		} finally {
			setIsLoading(false);
		}
	};

	const handleSignupFieldChange = (name: string, value: any) => {
		setSignupFormData((prev) => ({
			...prev,
			[name]: value,
		}));

		// Special case for email field to also update the login email
		if (name === "email") {
			setEmail(value);
		}
	};

	const renderAuthLink = (link: Link) => (
		<div className="text-sm">
			<p className="text-center space-x-2">
				{link.preText && <span>{link.preText} </span>}
				<a
					href={link.url}
					className="underline hover:text-foreground transition-colors"
				>
					{link.text}
				</a>
			</p>
		</div>
	);

	const renderResetPassword = () => (
		<Card className={cardClassName} style={cardStyle}>
			<CardHeader className={configCssClasses?.cardHeader}>
				<div className="flex items-center justify-center mb-4">
					{renderLogo()}
				</div>
				<CardTitle className={titleClasses} style={textStyle}>
					Reset Your Password
				</CardTitle>
				<CardDescription className={descriptionClasses} style={textStyle}>
					{resetPasswordToken
						? "Enter your new password below"
						: "Invalid or expired reset link. Please request a new password reset."}
				</CardDescription>
			</CardHeader>
			<CardContent className={configCssClasses?.cardContent}>
				{renderAlertMessages()}

				{resetPasswordToken ? (
					<form onSubmit={handleResetPassword}>
						<div className="grid gap-4">
							<div className="grid gap-2">
								<Label htmlFor="new-password" style={textStyle}>
									New Password
								</Label>
								<div className="relative">
									<Input
										id="new-password"
										type={showPassword ? "text" : "password"}
										value={newPassword}
										onChange={(e) => setNewPassword(e.target.value)}
										required
										className={inputClassName}
									/>
									<Button
										type="button"
										variant="ghost"
										size="icon"
										className="absolute right-0 top-0 h-full px-3"
										onClick={() => setShowPassword(!showPassword)}
									>
										{showPassword ? (
											<EyeOff className="h-4 w-4" />
										) : (
											<Eye className="h-4 w-4" />
										)}
									</Button>
								</div>
							</div>

							<div className="grid gap-2">
								<Label htmlFor="confirm-password" style={textStyle}>
									Confirm Password
								</Label>
								<div className="relative">
									<Input
										id="confirm-password"
										type={showPassword ? "text" : "password"}
										value={confirmPassword}
										onChange={(e) => setConfirmPassword(e.target.value)}
										required
										className={inputClassName}
									/>
								</div>
							</div>

							<Button
								type="submit"
								className={buttonClassName}
								disabled={isLoading}
							>
								{isLoading ? "Resetting Password..." : "Reset Password"}
							</Button>
						</div>
					</form>
				) : (
					<div className="grid gap-4">
						<Button
							variant="outline"
							type="button"
							onClick={() => setCurrentView("forgot-password")}
							className={configTheme.borderRadius}
						>
							Request New Reset Link
						</Button>
						<Button
							variant="outline"
							type="button"
							onClick={() => navigateToLogin()}
							className={configTheme.borderRadius}
						>
							<ArrowLeft className="mr-2 h-4 w-4" />
							Back to Login
						</Button>
					</div>
				)}
			</CardContent>
		</Card>
	);

	const renderFormFieldContent = (field: FormField) => {
		const { name, label, type, placeholder, required, options, validation } =
			field;
		const value = signupFormData[name];

		switch (type) {
			case "checkbox":
				return (
					<>
						<Checkbox
							id={`signup-${name}`}
							checked={value}
							onCheckedChange={(checked) =>
								handleSignupFieldChange(name, checked)
							}
						/>
						<Label htmlFor={`signup-${name}`}>{label}</Label>
					</>
				);

			case "select":
				return (
					<>
						<Label htmlFor={`signup-${name}`}>{label}</Label>
						<Select
							value={value}
							onValueChange={(newValue) =>
								handleSignupFieldChange(name, newValue)
							}
						>
							<SelectTrigger id={`signup-${name}`} className={inputClassName}>
								<SelectValue placeholder={placeholder || `Select ${label}`} />
							</SelectTrigger>
							<SelectContent>
								{options?.map((option) => (
									<SelectItem key={option.value} value={option.value}>
										{option.label}
									</SelectItem>
								))}
							</SelectContent>
						</Select>
					</>
				);

			case "password":
				return (
					<>
						<Label htmlFor={`signup-${name}`}>{label}</Label>
						<div className="relative">
							<Input
								id={`signup-${name}`}
								type={showPassword ? "text" : "password"}
								placeholder={placeholder}
								value={value}
								onChange={(e) => handleSignupFieldChange(name, e.target.value)}
								required={required}
								className={inputClassName}
								{...(validation && {
									pattern: validation.pattern,
									minLength: validation.minLength,
									maxLength: validation.maxLength,
								})}
							/>
							<Button
								type="button"
								variant="ghost"
								size="icon"
								className="absolute right-0 top-0 h-full px-3"
								onClick={() => setShowPassword(!showPassword)}
							>
								{showPassword ? (
									<EyeOff className="h-4 w-4" />
								) : (
									<Eye className="h-4 w-4" />
								)}
							</Button>
						</div>
					</>
				);

			default:
				return (
					<>
						<Label htmlFor={`signup-${name}`}>{label}</Label>
						<Input
							id={`signup-${name}`}
							type={type}
							placeholder={placeholder}
							value={value}
							onChange={(e) => handleSignupFieldChange(name, e.target.value)}
							required={required}
							className={inputClassName}
							{...(validation && {
								pattern: validation.pattern,
								minLength: validation.minLength,
								maxLength: validation.maxLength,
								min: validation.min,
								max: validation.max,
							})}
						/>
					</>
				);
		}
	};

	const renderFormField = (field: FormField) => {
		const { name, type, width = "full" } = field;

		// Define width classes
		const widthClasses = {
			full: "w-full",
			half: "w-full sm:w-[calc(50%-0.375rem)]", // Accounting for the gap
			third: "w-full sm:w-[calc(33.333%-0.5rem)]", // Accounting for the gap
		};

		// For checkbox type, use a different layout
		if (type === "checkbox") {
			return (
				<div
					className={`flex items-center space-x-2 ${widthClasses[width]}`}
					key={name}
				>
					{renderFormFieldContent(field)}
				</div>
			);
		}

		// For other field types
		return (
			<div className={`grid gap-2 ${widthClasses[width]}`} key={name}>
				{renderFormFieldContent(field)}
			</div>
		);
	};

	const renderSignup = () => {
		// Group fields by row
		const fieldsByRow: Record<string, FormField[]> = {};

		configSignupFields.forEach((field) => {
			const rowKey = field.row?.toString() || "default";
			if (!fieldsByRow[rowKey]) {
				fieldsByRow[rowKey] = [];
			}
			fieldsByRow[rowKey].push(field);
		});

		return (
			<form onSubmit={handleSignup}>
				<div className="grid gap-4">
					{Object.entries(fieldsByRow).map(([rowKey, fields]) => (
						<div
							key={rowKey}
							className="flex flex-wrap w-full gap-3 justify-between"
						>
							{fields.map((field) => renderFormField(field))}
						</div>
					))}

					<Button
						type="submit"
						className={buttonClassName}
						disabled={isLoading}
					>
						{isLoading ? "Creating Account..." : "Create Account"}
					</Button>

					{configLinks?.showAuthLinks &&
						configLinks.login &&
						renderAuthLink(configLinks.login)}

					{(configLinks.legal ?? []).length > 0 && (
						<div className="text-xs text-muted-foreground mt-2">
							<p className="text-center">
								{configLinks.legal?.map((link, index) => (
									<span key={index}>
										{index > 0 && " • "}
										<a
											href={link.url}
											className="underline hover:text-foreground transition-colors"
											target="_blank"
											rel="noopener noreferrer"
										>
											{link.text}
										</a>
									</span>
								))}
							</p>
						</div>
					)}
				</div>
			</form>
		);
	};

	const renderForgotPassword = () => (
		<Card className={cardClassName} style={cardStyle}>
			<CardHeader>
				<div className="flex items-center justify-center mb-4">
					{renderLogo()}
				</div>
				<CardTitle className={titleClasses} style={textStyle}>
					{configTitle}
				</CardTitle>
				<CardDescription className={descriptionClasses} style={textStyle}>
					Enter your email to receive a password reset link
				</CardDescription>
			</CardHeader>
			<CardContent>
				{renderAlertMessages()}
				<form onSubmit={handleForgotPassword}>
					<div className="grid gap-4">
						<div className="grid gap-2">
							<Label htmlFor="email-reset" style={textStyle}>
								Email
							</Label>
							<Input
								id="email-reset"
								type="email"
								placeholder="name@example.com"
								value={email}
								onChange={(e) => setEmail(e.target.value)}
								required
								className={inputClassName}
							/>
						</div>
						<Button
							type="submit"
							className={buttonClassName}
							disabled={isLoading}
						>
							{isLoading ? "Sending..." : "Send Reset Link"}
						</Button>
						<Button
							variant="outline"
							type="button"
							onClick={() => setCurrentView("login")}
							className={configTheme.borderRadius}
						>
							<ArrowLeft className="mr-2 h-4 w-4" />
							Back to Login
						</Button>
					</div>
				</form>
			</CardContent>
		</Card>
	);

	const renderVerifyOtp = () => (
		<Card className={cardClassName} style={cardStyle}>
			<CardHeader className={configCssClasses?.cardHeader}>
				<div className="flex items-center justify-center mb-4">
					{renderLogo()}
				</div>
				<CardTitle className={titleClasses} style={textStyle}>
					Verification Required
				</CardTitle>
				<CardDescription className={descriptionClasses} style={textStyle}>
					Enter the verification code sent to your device
				</CardDescription>
			</CardHeader>
			<CardContent className={configCssClasses?.cardContent}>
				<form onSubmit={handleVerifyOtp}>
					<div className="grid gap-4">
						<div
							className={cn({
								"grid gap-2": configVerification.input === "input",
								"flex justify-center": configVerification.input === "otp",
							})}
						>
							{configVerification.input === "input" && (
								<>
									<Label htmlFor="otp-code" style={textStyle}>
										Verification Code
									</Label>
									<Input
										id="otp-code"
										placeholder="Enter 6-digit code"
										value={otpCode}
										onChange={(e) => setOtpCode(e.target.value)}
										className={`text-center text-lg tracking-widest ${inputClassName}`}
										maxLength={codeInputLength.length}
									/>
								</>
							)}

							{configVerification.input === "otp" && (
								<InputOTP
									id="otp-code"
									maxLength={codeInputLength.length}
									pattern={REGEXP_ONLY_DIGITS_AND_CHARS}
									value={otpCode}
									required
									onChange={(v) => setOtpCode(v)}
									className={`text-center text-lg tracking-widest ${inputClassName}`}
								>
									<InputOTPGroup>
										{codeInputLength.map((value, idx) => (
											<InputOTPSlot index={idx} className="size-10 text-xl" />
										))}
									</InputOTPGroup>
								</InputOTP>
							)}
						</div>
						<Button
							type="submit"
							className={buttonClassName}
							disabled={isLoading}
						>
							{isLoading ? "Verifying..." : "Verify"}
						</Button>

						<div className="text-center mt-2">
							<Button
								variant="link"
								type="button"
								onClick={handleResendVerification}
								disabled={cooldownRemaining > 0 || isLoading}
								className="p-0 h-auto flex items-center justify-center mx-auto"
							>
								<RefreshCw className="mr-2 h-3 w-3" />
								{cooldownRemaining > 0
									? `Resend code (${cooldownRemaining}s)`
									: "Resend verification code"}
							</Button>
						</div>

						<Button
							variant="outline"
							type="button"
							onClick={() => {
								setCurrentView("login");
								resetState();
							}}
							className={configTheme.borderRadius}
						>
							<ArrowLeft className="mr-2 h-4 w-4" />
							Back to Login
						</Button>
					</div>
				</form>
			</CardContent>
		</Card>
	);

	const renderMfa = () => (
		<Card className={cardClassName} style={cardStyle}>
			<CardHeader>
				<div className="flex items-center justify-center mb-4">
					{renderLogo()}
				</div>
				<CardTitle className={titleClasses} style={textStyle}>
					Two-Factor Authentication
				</CardTitle>
				<CardDescription className={descriptionClasses} style={textStyle}>
					Enter the verification code sent to your device
				</CardDescription>
			</CardHeader>
			<CardContent>
				<form onSubmit={handleMfa}>
					<div className="grid gap-4">
						<div className="grid gap-2">
							<Label htmlFor="mfa-code" style={textStyle}>
								Verification Code
							</Label>
							<Input
								id="mfa-code"
								placeholder="Enter 6-digit code"
								value={mfaCode}
								onChange={(e) => setMfaCode(e.target.value)}
								className={`text-center text-lg tracking-widest ${inputClassName}`}
								maxLength={6}
							/>
						</div>
						<Button
							type="submit"
							className={buttonClassName}
							disabled={isLoading}
						>
							{isLoading ? "Verifying..." : "Verify"}
						</Button>
						<Button
							variant="outline"
							type="button"
							onClick={() => {
								setCurrentView("login");
								resetState();
							}}
							className={configTheme.borderRadius}
						>
							<ArrowLeft className="mr-2 h-4 w-4" />
							Back to Login
						</Button>
					</div>
				</form>
			</CardContent>
		</Card>
	);

	const renderLoginStep1 = () => (
		<form onSubmit={handleEmailSubmit}>
			<div className="grid gap-4">
				<div className="grid gap-2">
					<Label htmlFor="email" style={textStyle}>
						Email
					</Label>
					<Input
						id="email"
						type="email"
						placeholder="name@example.com"
						value={email}
						onChange={(e) => setEmail(e.target.value)}
						required
						className={inputClassName}
					/>
				</div>
				<Button type="submit" className={buttonClassName}>
					Continue
				</Button>
				{configLinks?.showAuthLinks &&
					configLinks.signup &&
					renderAuthLink(configLinks.signup)}

				{/*<div className="text-center">*/}
				{/*	<Button*/}
				{/*		variant="link"*/}
				{/*		type="button"*/}
				{/*		onClick={() => setCurrentView("forgot-password")}*/}
				{/*		className="p-0 h-auto"*/}
				{/*	>*/}
				{/*		Forgot password?*/}
				{/*	</Button>*/}
				{/*</div>*/}
			</div>
		</form>
	);

	const renderLoginStep2 = () => (
		<>
			<div className="mb-4">
				<Button
					variant="ghost"
					type="button"
					onClick={() => setLoginStep(1)}
					className="p-0 h-auto flex items-center text-muted-foreground hover:text-foreground"
				>
					<ArrowLeft className="mr-2 h-4 w-4" />
					{email}
				</Button>
			</div>

			{renderAlertMessages()}

			<div className="grid gap-4">
				{configSupportedMethods.includes("password") && (
					<form onSubmit={handleLogin}>
						<div className="grid gap-4">
							<div className="grid gap-2">
								<div className="flex items-center justify-between">
									<Label htmlFor="password" style={textStyle}>
										Password
									</Label>
									{configLinks?.forgotPassword ? (
										<a
											href={configLinks.forgotPassword.url}
											className="text-sm text-primary hover:underline"
										>
											{configLinks.forgotPassword.text}
										</a>
									) : (
										<Button
											variant="link"
											type="button"
											onClick={() => setCurrentView("forgot-password")}
											className="p-0 h-auto"
										>
											Forgot password?
										</Button>
									)}
								</div>
								<div className="relative">
									<Input
										id="password"
										type={showPassword ? "text" : "password"}
										value={password}
										onChange={(e) => setPassword(e.target.value)}
										required
										className={inputClassName}
									/>
									<Button
										type="button"
										variant="ghost"
										size="icon"
										className="absolute right-0 top-0 h-full px-3"
										onClick={() => setShowPassword(!showPassword)}
									>
										{showPassword ? (
											<EyeOff className="h-4 w-4" />
										) : (
											<Eye className="h-4 w-4" />
										)}
									</Button>
								</div>
							</div>
							<Button
								type="submit"
								className={buttonClassName}
								disabled={isLoading}
							>
								{isLoading ? "Signing in..." : "Sign in"}
							</Button>
						</div>
					</form>
				)}

				{configSupportedMethods.includes("passwordless") && (
					<Button
						variant="outline"
						onClick={handlePasswordless}
						disabled={isLoading}
						className={configTheme.borderRadius}
					>
						<Mail className="mr-2 h-4 w-4" />
						{isLoading ? "Sending..." : "Sign in with Email Link"}
					</Button>
				)}

				{configSupportedMethods.includes("passkey") && (
					<Button
						variant="outline"
						onClick={handlePasskey}
						disabled={isLoading}
						className={configTheme.borderRadius}
					>
						<Fingerprint className="mr-2 h-4 w-4" />
						{isLoading ? "Authenticating..." : "Sign in with Passkey"}
					</Button>
				)}
			</div>

			{configOauthProviders.length > 0 && (
				<>
					<div className="relative my-4">
						<div className="absolute inset-0 flex items-center">
							<Separator />
						</div>
						<div className="relative flex justify-center text-xs uppercase">
							<span
								className="bg-background px-2 text-muted-foreground"
								style={textStyle}
							>
								Or continue with
							</span>
						</div>
					</div>

					<div className="grid gap-2">
						{configOauthProviders.map((provider, index) => (
							<Button
								key={index}
								variant="outline"
								onClick={provider.onClick}
								className={configTheme.borderRadius}
							>
								{provider.icon}
								<span className="ml-2">{provider.name}</span>
							</Button>
						))}
					</div>
				</>
			)}
		</>
	);

	// Render different views based on currentView state
	if (currentView === "forgot-password") {
		return renderForgotPassword();
	}

	if (currentView === "verify-otp") {
		return renderVerifyOtp();
	}

	if (currentView === "mfa") {
		return renderMfa();
	}

	return (
		<Card className={cardClassName} style={cardStyle}>
			<CardHeader>
				<div
					className={cn("flex mb-4 items-center", {
						"justify-center": configTitleAlign === "center",
						"justify-start": configTitleAlign === "left",
						"justify-end": configTitleAlign === "right",
					})}
				>
					{renderLogo()}
				</div>
				<CardTitle className={titleClasses} style={textStyle}>
					{configTitle}
				</CardTitle>
				<CardDescription className={descriptionClasses} style={textStyle}>
					{configDescription}
				</CardDescription>
			</CardHeader>
			<CardContent>
				{configShowTabs && configAvailableTabs.length > 1 ? (
					<Tabs
						defaultValue={activeTab}
						value={activeTab}
						onValueChange={handleTabChange}
					>
						<TabsList
							className={`grid w-full ${configTheme.borderRadius}`}
							style={{
								gridTemplateColumns: `repeat(${configAvailableTabs.length}, 1fr)`,
							}}
						>
							{configAvailableTabs.includes("login") && (
								<TabsTrigger value="login">Login</TabsTrigger>
							)}
							{configAvailableTabs.includes("signup") && (
								<TabsTrigger value="signup">Sign Up</TabsTrigger>
							)}
						</TabsList>
						{configAvailableTabs.includes("login") && (
							<TabsContent value="login" className="mt-4">
								{loginStep === 1 ? renderLoginStep1() : renderLoginStep2()}
							</TabsContent>
						)}
						{configAvailableTabs.includes("signup") && (
							<TabsContent value="signup" className="mt-4">
								{renderSignup()}
							</TabsContent>
						)}
					</Tabs>
				) : (
					<div className="mt-4">
						{activeTab === "login"
							? loginStep === 1
								? renderLoginStep1()
								: renderLoginStep2()
							: renderSignup()}
					</div>
				)}
			</CardContent>
		</Card>
	);
}
