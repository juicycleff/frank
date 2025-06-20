import {
    AuthenticationApi,
    ChangePasswordRequest,
    Configuration,
    MFAApi,
    MFABackCodes,
    MFASetupResponse,
    MFASetupVerifyResponse,
    PaginatedOutputMFAMethod,
    PaginatedOutputPasskeySummary,
    PasskeyRegistrationBeginRequest,
    PasskeyRegistrationBeginResponse,
    PasskeyRegistrationFinishRequest,
    PasskeyRegistrationFinishResponse,
    PasskeysApi,
    PasskeySummary,
    ResendVerificationResponse,
    SetupMFARequest,
    UpdatePasskeyRequest,
    User,
    UserProfileUpdateRequest,
    UsersApi,
    VerifyMFASetupRequest,
} from '@frank-auth/client';

import {FrankAuthConfig, FrankAuthError} from './index';
import {handleError} from "./errors";

export class FrankUser {
    private config: FrankAuthConfig;
    private usersApi: UsersApi;
    private authApi: AuthenticationApi;
    private mfaApi: MFAApi;
    private passkeyApi: PasskeysApi;
    private accessToken: string | null = null;

    constructor(config: FrankAuthConfig, accessToken?: string) {
        this.config = config;
        this.accessToken = accessToken || null;

        const configuration = new Configuration({
            basePath: config.apiUrl,
            accessToken: () => this.accessToken || '',
            credentials: 'include',
            headers: {
                'X-Publishable-Key': config.publishableKey,
            },
        });

        this.usersApi = new UsersApi(configuration);
        this.authApi = new AuthenticationApi(configuration);
        this.mfaApi = new MFAApi(configuration);
        this.passkeyApi = new PasskeysApi(configuration);
    }

    // Update access token (called by FrankAuth when token changes)
    setAccessToken(token: string | null): void {
        this.accessToken = token;
    }

    // Profile management
    async getProfile(): Promise<User> {
        try {
            return await this.usersApi.getUserProfile();
        } catch (error) {
            throw await handleError(error)
        }
    }

    async updateProfile(request: UserProfileUpdateRequest): Promise<User> {
        try {
            return await this.usersApi.updateUserProfile({ userProfileUpdateRequest: request });
        } catch (error) {
            throw await handleError(error)
        }
    }

    async changePassword(request: ChangePasswordRequest): Promise<void> {
        try {
            await this.usersApi.changePassword({ changePasswordRequest: request });
        } catch (error) {
            throw await handleError(error)
        }
    }

    // Email verification
    async resendEmailVerification(email?: string): Promise<ResendVerificationResponse> {
        try {
            return await this.authApi.resendVerification({
                resendVerificationRequest: { email, type: 'email' }
            });
        } catch (error) {
            throw await handleError(error)
        }
    }

    async resendPhoneVerification(phone?: string): Promise<ResendVerificationResponse> {
        try {
            return await this.authApi.resendVerification({
                resendVerificationRequest: { phoneNumber: phone, type: 'sms' }
            });
        } catch (error) {
            throw await handleError(error)
        }
    }

    // MFA management
    async getMFAMethods(options?: {
        limit?: number;
        offset?: number;
    }): Promise<PaginatedOutputMFAMethod> {
        try {
            return await this.mfaApi.listMFAMethods({
                orgId: '',
                userId: '',
            });
        } catch (error) {
            throw await handleError(error)
        }
    }

    async setupMFA(request: SetupMFARequest): Promise<MFASetupResponse> {
        try {
            return await this.authApi.setupMFA({ setupMFARequest: request });
        } catch (error) {
            throw await handleError(error)
        }
    }

    async verifyMFASetup(request: VerifyMFASetupRequest): Promise<MFASetupVerifyResponse> {
        try {
            return await this.authApi.verifyMFASetup({ verifyMFASetupRequest: request });
        } catch (error) {
            throw await handleError(error)
        }
    }

    async disableMFA(): Promise<void> {
        try {
            await this.authApi.disableMFA();
        } catch (error) {
            throw await handleError(error)
        }
    }

    async getBackupCodes(regenerate = false): Promise<MFABackCodes> {
        try {
            return await this.authApi.getMFABackupCodes({
                generateBackupCodesRequest: { count: regenerate ? 1 : undefined }
            });
        } catch (error) {
            throw await handleError(error)
        }
    }

    // Passkey management
    async getPasskeys(options?: {
        limit?: number;
        offset?: number;
        search?: string;
    }): Promise<PaginatedOutputPasskeySummary> {
        try {
            return await this.authApi.listPasskeys({
                limit: options?.limit,
                offset: options?.offset,
                search: options?.search,
            });
        } catch (error) {
            throw await handleError(error)
        }
    }

    async createPasskey(request: PasskeyRegistrationBeginRequest): Promise<{
        beginResponse: PasskeyRegistrationBeginResponse;
        finishRegistration: (finishRequest: PasskeyRegistrationFinishRequest) => Promise<PasskeyRegistrationFinishResponse>;
    }> {
        try {
            const beginResponse = await this.passkeyApi.beginPasskeyRegistration({
                passkeyRegistrationBeginRequest: request
            });

            return {
                beginResponse,
                finishRegistration: async (finishRequest: PasskeyRegistrationFinishRequest) => {
                    return await this.passkeyApi.finishPasskeyRegistration({
                        passkeyRegistrationFinishRequest: finishRequest
                    });
                }
            };
        } catch (error) {
            throw await handleError(error)
        }
    }

    async deletePasskey(passkeyId: string): Promise<void> {
        try {
            await this.authApi.deletePasskey({ id: passkeyId });
        } catch (error) {
            throw await handleError(error)
        }
    }

    async updatePasskey(passkeyId: string, request: UpdatePasskeyRequest): Promise<PasskeySummary> {
        try {
            // Note: This endpoint might not exist in the current API
            // For now, we'll throw an error or implement a workaround
            throw new FrankAuthError('Update passkey not implemented', 'NOT_IMPLEMENTED');
        } catch (error) {
            throw await handleError(error)
        }
    }

    // User statistics and insights
    async getUserStats(): Promise<{
        totalSessions: number;
        totalPasskeys: number;
        mfaEnabled: boolean;
        lastLoginAt?: Date;
        accountCreatedAt?: Date;
        emailVerified: boolean;
        phoneVerified: boolean;
    }> {
        try {
            const [profile, sessions, passkeys] = await Promise.all([
                this.getProfile(),
                // Note: We'd need to import FrankSession or use the sessions API directly
                Promise.resolve({ data: [] }), // Placeholder
                this.getPasskeys(),
            ]);

            return {
                totalSessions: sessions.data?.length || 0,
                totalPasskeys: passkeys.pagination?.totalPages || 0,
                mfaEnabled: profile.mfaEnabled || false,
                // lastLoginAt: profile.lastLoginAt ? new Date(profile.lastLoginAt) : undefined,
                accountCreatedAt: profile.createdAt ? new Date(profile.createdAt) : undefined,
                emailVerified: profile.emailVerified || false,
                phoneVerified: profile.phoneVerified || false,
            };
        } catch (error) {
            throw await handleError(error)
        }
    }

    // Security utilities
    async getSecuritySummary(): Promise<{
        securityScore: number;
        recommendations: string[];
        risks: string[];
        mfaEnabled: boolean;
        passkeyCount: number;
        recentSuspiciousActivity: boolean;
    }> {
        try {
            const [profile, passkeys] = await Promise.all([
                this.getProfile(),
                this.getPasskeys(),
            ]);

            const recommendations: string[] = [];
            const risks: string[] = [];
            let securityScore = 100;

            // Check MFA
            if (!profile.mfaEnabled) {
                recommendations.push('Enable multi-factor authentication for better security');
                risks.push('Account not protected by MFA');
                securityScore -= 30;
            }

            // Check passkeys
            const passkeyCount = passkeys.pagination?.total || 0;
            if (passkeyCount === 0) {
                recommendations.push('Set up passkeys for passwordless authentication');
                securityScore -= 20;
            }

            // Check email verification
            if (!profile.emailVerified) {
                recommendations.push('Verify your email address');
                risks.push('Email address not verified');
                securityScore -= 15;
            }

            // Check password strength (if applicable)
            if (!profile.passwordless && !profile.mfaEnabled) {
                risks.push('Account relies only on password authentication');
                securityScore -= 20;
            }

            return {
                securityScore: Math.max(0, securityScore),
                recommendations,
                risks,
                mfaEnabled: profile.mfaEnabled || false,
                passkeyCount,
                recentSuspiciousActivity: false, // This would require session analysis
            };
        } catch (error) {
            throw await handleError(error)
        }
    }

    // Profile validation
    validateProfile(profile: Partial<UserProfileUpdateRequest>): {
        isValid: boolean;
        errors: Record<string, string[]>;
    } {
        const errors: Record<string, string[]> = {};

        if (profile.email && !this.isValidEmail(profile.email)) {
            errors.email = ['Please provide a valid email address'];
        }

        if (profile.phone && !this.isValidPhone(profile.phone)) {
            errors.phone = ['Please provide a valid phone number'];
        }

        if (profile.firstName && profile.firstName.length < 1) {
            errors.firstName = ['First name is required'];
        }

        if (profile.lastName && profile.lastName.length < 1) {
            errors.lastName = ['Last name is required'];
        }

        return {
            isValid: Object.keys(errors).length === 0,
            errors,
        };
    }

    // Utility methods
    private isValidEmail(email: string): boolean {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
    }

    private isValidPhone(phone: string): boolean {
        const phoneRegex = /^\+?[\d\s\-\(\)]+$/;
        return phoneRegex.test(phone) && phone.replace(/\D/g, '').length >= 10;
    }
}