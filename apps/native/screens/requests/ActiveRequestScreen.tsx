import { useNavigation } from "@react-navigation/native";
import {
  SafeAreaView,
  View,
  TouchableOpacity,
  Text,
  Image,
} from "react-native";
import { Ionicons, Feather } from "@expo/vector-icons";
import { useEffect } from "react";
import { fetchActiveJobs } from "@api/plugs-api";
import { useDispatch, useSelector } from "react-redux";
import { Dispatch, RootState } from "../../store";
import { ActiveJobPlugView } from "../ActiveJobPlugView";
import { TrackingItem } from "../../components/tracking/TrackingItem";

// TODO: Figure this shit out
// is_active | status | request_status?
//                      - accepted
//                      - in_transit

// Active request status
//  Current status
// Request status
//  Overall status
// All request stauses
//  Info on all statuses and if they're true/false
//

export const ActiveRequestScreen = () => {
  const navigation = useNavigation();
  const dispatch = useDispatch<Dispatch>();
  const { activeJob } = useSelector((state: RootState) => state.jobs);
  const { role } = useSelector((state: RootState) => state.auth);

  useEffect(() => {
    fetchActiveJobs().then((data) => dispatch.jobs.populateActiveJob(data));
    console.log("populate active job");
  }, [dispatch.jobs]);

  if (role === "plug") return <ActiveJobPlugView />;

  return (
    <SafeAreaView className="flex-1 bg-white relative">
      <View className="px-4 flex flex-row items-center space-x-4">
        <TouchableOpacity
          onPress={navigation.goBack}
          className="bg-gray-100 rounded-full px-4 py-1"
        >
          <Feather name="arrow-left" size={18} color="#1f2937" />
        </TouchableOpacity>
        <Text className="text-2xl font-medium text-black">
          {activeJob?.title}
        </Text>
      </View>

      <View className="px-4 mt-4">
        <Text className="text-gray-500 text-lg">{activeJob?.description}</Text>

        <Text>Current status: {activeJob.status}</Text>

        {activeJob.status && activeJob.tracking_status && (
          <View className="mt-8 space-y-4">
            <View>
              <TrackingItem
                label={`@${activeJob.username} har godatt jobben.`}
                active={activeJob.status === "accepted"}
                completed={activeJob.tracking_status["accepted"] === true}
              />
            </View>

            <View>
              <TrackingItem
                label="Plugen er på vei."
                active={activeJob.status === "in_transit"}
                completed={activeJob.tracking_status["in_transit"] === true}
              />
            </View>

            <View>
              <TrackingItem
                label="Jobben utføres."
                active={activeJob.status === "active"}
                completed={activeJob.tracking_status["active"] === true}
              />
            </View>

            <View>
              <TrackingItem
                label="Jobben er utført og betalt."
                active={activeJob.status === "completed"}
                completed={activeJob.tracking_status["completed"] === true}
              />
            </View>
          </View>
        )}
      </View>

      <View className="mt-4 absolute bottom-10 px-4 left-0 right-0">
        <Text className="font-medium text-xl text-black">Plug</Text>
        <View className="flex-row justify-between items-center">
          <View className="flex flex-row items-center space-x-2 mt-2">
            <Image
              source={{ uri: activeJob?.avatar }}
              className="h-10 w-10 rounded-full"
            />
            <Text className="text-lg font-medium underline text-gray-500">
              @{activeJob?.username}
            </Text>
          </View>

          <TouchableOpacity className="bg-rose-100 h-10 w-10 flex items-center justify-center rounded-lg">
            <Ionicons name="chatbubble" color="#9f1239" size={24} />
          </TouchableOpacity>
        </View>
      </View>
    </SafeAreaView>
  );
};
